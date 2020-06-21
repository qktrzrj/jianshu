package resolve

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/plugins/sqlog"
)

type likeResolver struct{}

var LikeResolver likeResolver

type objArg struct {
	Id      int           `graphql:"id"`
	ObjType model.Objtype `graphql:"objType"`
}

// 是否已点赞
func (r likeResolver) HasLike(ctx context.Context, args objArg) (bool, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	hasLike, err := model.HasLike(tx, userId, args.Id, args.ObjType)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return false, err
	}
	return hasLike, nil
}

// 点赞
func (r likeResolver) Like(ctx context.Context, args objArg) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	err := model.Like(tx, map[string]interface{}{
		"objtype": args.ObjType,
		"objid":   args.Id,
		"uid":     userId,
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return errors.New("点赞失败")
	}

	if args.ObjType == model.ArticleObj {
		err := model.AddViewOrLikeOrCmt(tx, args.Id, 1, true)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return errors.New("点赞失败")
		}
		go ArticleResolver.PutHots(ctx, args.Id)
	}
	return nil
}

// 取消点赞
func (r likeResolver) UnLike(ctx context.Context, args objArg) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	err := model.UnLike(tx, args.ObjType, args.Id, userId)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return errors.New("取消点赞失败")
	}
	if args.ObjType == model.ArticleObj {
		err := model.AddViewOrLikeOrCmt(tx, args.Id, 1, false)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return errors.New("取消点赞失败")
		}
		go ArticleResolver.PutHots(ctx, args.Id)
	}
	return nil
}
