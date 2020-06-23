package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql"
	"github.com/shyptr/graphql/errors"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/util"
	"github.com/shyptr/plugins/sqlog"
	"net/http"
	"time"
)

func CORS() graphql.HandlerFunc {
	return func(c *graphql.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "http://www.shyptr.cn")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,X-Requested-With")
		if c.Request.Method == http.MethodOptions {
			return
		}
		c.Next()
	}
}

func Logger() graphql.HandlerFunc {
	return func(c *graphql.Context) {
		logger := util.GetLogger()
		c.Set("logger", logger)
		start := time.Now()
		defer func() {
			reqMethod := c.Method
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			operationName := c.OperationName
			if operationName == "" {
				operationName = "query"
			}
			logger.Info().Int("status", statusCode).Interface("method", reqMethod).TimeDiff("latencyTime", start, time.Now()).
				Str("ip", clientIP).Interface("operationName", operationName).Send()
			util.PutLogger(logger)
		}()
		c.Next()
	}
}

func Recovery() graphql.HandlerFunc {
	return func(c *graphql.Context) {
		logger := c.Value("logger").(zerolog.Logger)
		defer func() {
			if r := recover(); r != nil {
				logger.Error().Interface("[Recovery] panic received", r).Send()
				c.Error = append(c.Error, errors.New("%v", r))
				c.ServerError("", http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func Tx(ld *sqlog.DB) graphql.HandlerFunc {
	db := ld.Runner.(*sqlx.DB)
	return func(c *graphql.Context) {
		logger := c.Value("logger").(zerolog.Logger)
		tx, err := db.Beginx()
		if err != nil {
			logger.Error().AnErr("事务开启失败", err).Send()
			c.ServerError("", http.StatusInternalServerError)
			return
		}
		c.Set("tx", &sqlog.DB{Runner: tx, Logger: ld.Logger})
		defer func() {
			if c.Err() != nil {
				tx.Rollback()
				return
			}
			if err := tx.Commit(); err != nil {
				logger.Error().AnErr("transition commit failed", err).Send()
				tx.Rollback()
			}
		}()
		c.Next()
	}
}

func BasicAuth() graphql.HandlerFunc {
	return func(c *graphql.Context) {
		logger := c.Value("logger").(zerolog.Logger)
		// sessionId
		cs, err := c.Request.Cookie("Sess")
		if err == http.ErrNoCookie || cs == nil {
			sessionId := uuid.New().String()
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    "Sess",
				Value:   sessionId,
				Path:    "/",
				Expires: time.Now().Add(time.Hour * 24),
				MaxAge:  int(time.Now().Add(time.Hour * 24).Unix()),
				Domain:  "www.shyptr.cn",
			})
			c.Set("sessionId", sessionId)
		} else {
			c.Set("sessionId", cs.Value)
		}

		cookie, err := c.Request.Cookie("me")
		defer c.Next()
		if err == http.ErrNoCookie || cookie == nil {
			return
		}
		token := cookie.Value
		if token != "" {
			user, err := util.ParseToken(token)
			if err != nil {
				if err == util.ErrTokenExpire {
					c.Set("sessionId", token)
					return
				} else {
					logger.Error().AnErr("解析token失败", err).Send()
					return
				}
			}
			c.Set("userId", user.Id)
			c.Set("root", user.Root)
			c.Set("userState", user.State)
		}
	}
}

func LoginNeed() schemabuilder.ExecuteFunc {
	return func(ctx context.Context, args, source interface{}) error {
		id := ctx.Value("userId")
		if id == nil {
			c := ctx.(*graphql.Context)
			c.ServerError("", http.StatusUnauthorized)
			return errors.New("你必须先登录")
		}
		return nil
	}
}

func NotLogin() schemabuilder.ExecuteFunc {
	return func(ctx context.Context, args, source interface{}) error {
		id := ctx.Value("userId")
		if id != nil {
			c := ctx.(*graphql.Context)
			c.ServerError("", http.StatusMethodNotAllowed)
			return errors.New("你必须先退出登录")
		}
		return nil
	}
}
