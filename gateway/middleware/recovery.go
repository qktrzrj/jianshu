package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net"
	"net/http/httputil"
	"os"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := ctx.Value("logger").(zerolog.Logger)
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				if brokenPipe {
					logger.Error().Interface("error", err).Interface("request", httpRequest).Send()
				} else {
					logger.Error().Caller(3).Interface("error", err).Msgf("[Recovery] %s panic recovered.", strings.Join(headers, "\r\n"))
				}
			}
		}()
		ctx.Next()
	}
}
