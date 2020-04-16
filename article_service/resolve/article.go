package resolve

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shyptr/hello-world-web/article_service/model"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
)

type articleResolve struct{}

var ArticleResolve = articleResolve{}

func (this articleResolve) ArticleCount(ctx context.Context, article model.Article) (model.ArticleEx, error) {
	return model.GetArticleEx(ctx, article.Id)
}

func (this articleResolve) Article(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) (model.Article, error) {
	return model.GetArticle(ctx, sqlex.Eq{"id": args.Id})
}

func (this articleResolve) Articles(ctx context.Context, args struct {
	Title   *string  `graphql:"title"`
	Uid     *int64   `graphql:"uid"`
	Content *string  `graphql:"content"`
	Tags    []string `graphql:"tags"`
}) ([]model.Article, error) {
	return model.GetArticles(ctx,
		sqlex.IF{Condition: args.Title != nil, Sq: sqlex.Like{"title": args.Title}},
		sqlex.IF{Condition: args.Uid != nil, Sq: sqlex.Eq{"uid": args.Uid}},
		sqlex.IF{Condition: args.Content != nil, Sq: sqlex.Like{"content": args.Content}},
		sqlex.IF{Condition: args.Tags != nil, Sq: sqlex.Eq{"tags": args.Tags}},
	)
}

func (this articleResolve) AddArticle(ctx context.Context, args struct {
	Title   string   `graphql:"title" validate:"max=50"`
	Cover   *string  `graphql:"cover"`
	Content string   `graphql:"content"`
	Tags    []string `graphql:"tags"`
}) (model.Article, error) {
	return model.InsertArticle(ctx, map[string]interface{}{
		"uid":     ctx.Value("userId"),
		"title":   args.Title,
		"cover":   args.Cover,
		"content": args.Content,
		"tags":    args.Tags,
	})
}

func (this articleResolve) UpdateArticle(ctx context.Context, args struct {
	Id      int64    `graphql:"id"`
	Title   string   `graphql:"title" validate:"max=50"`
	Cover   *string  `graphql:"cover"`
	Content string   `graphql:"content"`
	Tags    []string `graphql:"tags"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	row := model.PSql.Select("count(id)").From("article").WhereExpr(
		sqlex.Eq{"uid": ctx.Value("userId")},
		sqlex.Eq{"id": args.Id},
	).RunWith(tx).QueryRow()
	var count int
	err := row.Scan(&count)
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("validate article owner failed")
	}
	if count == 0 {
		return fmt.Errorf("there are no artilce %s for user %s", args.Id, ctx.Value("userId"))
	}

	return model.UpdateArticle(ctx, map[string]interface{}{
		"title":   args.Title,
		"cover":   args.Cover,
		"content": args.Content,
		"tags":    args.Tags,
	}, args.Id)
}

func (this articleResolve) DeleteArticle(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	row := model.PSql.Select("count(id)").From("article").WhereExpr(
		sqlex.Eq{"uid": ctx.Value("userId")},
		sqlex.Eq{"id": args.Id},
	).RunWith(tx).QueryRow()
	var count int
	err := row.Scan(&count)
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("validate article owner failed")
	}
	if count == 0 {
		return fmt.Errorf("there are no artilce %s for user %s", args.Id, ctx.Value("userId"))
	}
	return model.RemoveArticle(ctx, args.Id)
}
