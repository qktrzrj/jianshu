package resolve

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/cache"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/plugins/sqlog"
)

type msgResolver struct{}

var MsgResolver msgResolver

func (r msgResolver) MsgNum(ctx context.Context) (model.MsgNum, error) {
	userId := ctx.Value("userId").(int)

	msgNum, err := cache.QueryCache(ctx, cache.MsgNum{Uid: userId}, func() (interface{}, error) {
		return model.QueryMsgNum(ctx.Value("tx").(*sqlog.DB), userId)
	})
	if err != nil {
		logger := ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller().Err(err).Send()
		return model.MsgNum{}, err
	}
	return msgNum.(model.MsgNum), nil
}

func (r msgResolver) ListMsg(ctx context.Context, args struct {
	Typ model.MsgType `graphql:"typ"`
}) ([]model.Msg, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	userId := ctx.Value("userId").(int)

	msgs, err := model.ListMsg(tx, args.Typ, userId)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, err
	}
	cache.Delete(cache.MsgNum{Uid: userId}.GetCacheKey())
	err = model.ReadMsg(tx, ctx.Value("userId").(int), args.Typ)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, err
	}
	return msgs, nil
}

func (r msgResolver) AddMsg(ctx context.Context, args struct {
	FromId  int           `graphql:"fromId"`
	ToId    int           `graphql:"toId"`
	Content string        `graphql:"content"`
	Typ     model.MsgType `graphql:"typ"`
}) error {
	cache.Delete(cache.MsgNum{Uid: args.ToId}.GetCacheKey())
	cache.Delete(cache.Msg{Uid: args.ToId, Typ: args.Typ}.GetCacheKey())
	err := model.AddMsg(ctx.Value("tx").(*sqlog.DB), args.Typ, args.FromId, args.ToId, args.Content)
	if err != nil {
		logger := ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller().Err(err).Send()
		return err
	}
	return nil
}
