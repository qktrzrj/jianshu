package resolve

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/plugins/sqlog"
)

type replyResolver struct{}

var ReplyResolver replyResolver

func (r replyResolver) List(ctx context.Context, args IdArgs) ([]model.Reply, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	replies, err := model.ListReply(tx, args.Id)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, errors.New("获取评论回复失败")
	}
	return replies, nil
}

func (r replyResolver) Add(ctx context.Context, args struct {
	Id      int    `graphql:"id"`
	Content string `graphql:"content"`
}) (model.Reply, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	reply, err := model.AddReply(tx, map[string]interface{}{
		"cid":     args.Id,
		"uid":     userId,
		"content": args.Content,
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.Reply{}, err
	}
	return reply, nil
}
