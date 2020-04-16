package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/hello-world-web/util"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		logger := util.NewLogger()
		ctx.Set("logger", logger)
		defer func() {
			statusCode := ctx.Writer.Status()
			clientIP := ctx.ClientIP()
			method := ctx.Value("method")
			if method == nil || method == "" {
				method = "query"
			}
			logger.Info().Int("status", statusCode).TimeDiff("latencyTime", time.Now(), startTime).
				Str("ip", clientIP).Interface("method", method).Interface("operationName", ctx.Value("operationName")).Send()
			util.PutLogger(logger)
		}()
		ctx.Next()
	}
}
