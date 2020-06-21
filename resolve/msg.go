package resolve

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/plugins/sqlog"
)

type msgResolver struct{}

var MsgResolver msgResolver

func (r msgResolver) MsgNum(ctx context.Context) (model.MsgNum, error) {
	msgNum, err := model.QueryMsgNum(ctx.Value("tx").(*sqlog.DB), ctx.Value("userId").(int))
	if err != nil {
		logger := ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller().Err(err).Send()
		return model.MsgNum{}, err
	}
	return msgNum, nil
}

func (r msgResolver) ListMsg(ctx context.Context, args struct {
	Typ model.MsgType `graphql:"typ"`
}) ([]model.Msg, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	msgs, err := model.ListMsg(tx, args.Typ, ctx.Value("userId").(int))
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, err
	}
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
	err := model.AddMsg(ctx.Value("tx").(*sqlog.DB), args.Typ, args.FromId, args.ToId, args.Content)
	if err != nil {
		logger := ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller().Err(err).Send()
		return err
	}
	return nil
}
