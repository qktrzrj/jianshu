package middleware

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"google.golang.org/grpc"
)

func Tx(sl *sqlog.DB) grpc.UnaryServerInterceptor {
	db := sl.Runner.(*sqlx.DB)
	return func(ctx context.Context, r interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (resp interface{}, err error) {
		tx, err := db.Beginx()
		logger := ctx.Value("logger").(zerolog.Logger)
		if err != nil {
			logger.Error().Err(err).Msg("transition begin error")
			return
		}
		ctx = context.WithValue(ctx, "tx", &sqlog.DB{Runner: tx, Logger: sl.Logger})
		defer func() {
			if ctx.Err() != nil || err != nil {
				tx.Rollback()
				return
			}
			if e := tx.Commit(); e != nil {
				tx.Rollback()
			}
		}()
		return h(ctx, r)
	}
}
