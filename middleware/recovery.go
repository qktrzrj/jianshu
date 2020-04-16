package middleware

import (
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, r interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (resp interface{}, err error) {
		logger := ctx.Value("logger").(zerolog.Logger)
		defer func() {
			if e := recover(); e != nil {
				logger.Error().Caller(3).Interface("error", e).Msg("[Recovery] %s panic recovered.")
				err = status.Errorf(codes.Internal, "Panic err: %v", e)
			}
		}()
		return h(ctx, r)
	}
}
