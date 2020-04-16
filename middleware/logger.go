package middleware

import (
	"context"
	"github.com/shyptr/hello-world-web/util"
	"google.golang.org/grpc"
	"time"
)

func Logger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, r interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		logger := util.NewLogger()
		ctx = context.WithValue(ctx, "logger", logger)
		defer func() {
			logger.Info().TimeDiff("latencyTime", time.Now(), startTime).Interface("method", info.FullMethod).Send()
			util.PutLogger(logger)
		}()
		return h(ctx, r)
	}
}
