package resolve

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/plugins/sqlog"
)

type commentResolver struct{}

var CommentResolver commentResolver

func (c commentResolver) List(ctx context.Context, article model.Article) ([]model.Comment, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	commentList, err := model.CommentList(tx, article.Id)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, errors.New("获取文章评论失败")
	}
	return commentList, nil
}

func (c commentResolver) Add(ctx context.Context, args struct {
	Id      int    `graphql:"id"`
	Content string `graphql:"content"`
}) (model.Comment, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	comment, err := model.AddComment(tx, map[string]interface{}{
		"aid":     args.Id,
		"uid":     userId,
		"content": args.Content,
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.Comment{}, err
	}
	err = model.AddViewOrLikeOrCmt(tx, args.Id, 2, true)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.Comment{}, err
	}
	return comment, nil
}
