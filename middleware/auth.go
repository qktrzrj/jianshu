package middleware

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/hello-world-web/util"
	"strconv"
)

func Auth() schemabuilder.ExecuteFunc {
	return func(ctx context.Context, args, source interface{}) error {
		token, ok := metadata.Get(ctx, "me")
		if ok && token != "" {
			validateToken, err := util.ParseAnValidateToken(token)
			if err == nil {
				id, _ := strconv.Atoi(validateToken.Id)
				ctx = context.WithValue(ctx, "userId", int64(id))
			}
		}
		return nil
	}
}

func LoginNeed() schemabuilder.ExecuteFunc {
	return func(ctx context.Context, args, source interface{}) error {
		token, ok := metadata.Get(ctx, "me")
		if ok && token != "" {
			validateToken, err := util.ParseAnValidateToken(token)
			if err == nil {
				id, _ := strconv.Atoi(validateToken.Id)
				ctx = context.WithValue(ctx, "userId", int64(id))
				return nil
			}
		}
		return fmt.Errorf("must login")
	}
}

func NotLogin() schemabuilder.ExecuteFunc {
	return func(ctx context.Context, args, source interface{}) error {
		logger := ctx.Value("logger").(zerolog.Logger)
		logger.Info().Str("GetToken", "================").Send()
		token, ok := metadata.Get(ctx, "me")
		if ok && token != "" {
			logger.Info().Str("TokenValue", token).Send()
			_, err := util.ParseAnValidateToken(token)
			if err == nil {
				return fmt.Errorf("must logout")
			}
		}
		return nil
	}
}
