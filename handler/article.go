package handler

import (
	"context"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/resolve"
)

func registerArticle(schema *schemabuilder.Schema) {
	// 枚举
	schema.Enum("ArticleState", model.ArticleState(0), map[string]model.ArticleState{
		"Unaudited": model.Unaudited,
		"Online":    model.Online,
		"Offline":   model.Offline,
		"Deleted":   model.Deleted,
	})

	article := schema.Object("Article", model.Article{})
	// 文章扩展字段:浏览数/评论数/点赞数
	article.FieldFunc("ViewNum", func(source model.Article) int { return source.ViewNum })
	article.FieldFunc("CmtNum", func(source model.Article) int { return source.CmtNum })
	article.FieldFunc("LikeNum", func(source model.Article) int { return source.LikeNum })
	// 文章作者
	article.FieldFunc("User", func(ctx context.Context, source model.Article) (model.User, error) {
		return resolve.UserResolver.User(ctx, resolve.IdArgs{Id: source.Uid})
	})

	query := schema.Query()
	schemabuilder.RelayKey(model.Article{}, "id")
	schema.Object("articlePage", resolve.ArticlePage{})
	// 热门文章（分页）
	query.FieldFunc("HotArticles", resolve.ArticleResolver.Hots, schemabuilder.RelayConnection)
	// 查询文章（分页）
	query.FieldFunc("Articles", resolve.ArticleResolver.Articles, schemabuilder.RelayConnection)
	// 获取登录人文章（分页）
	query.FieldFunc("CurArticles", resolve.ArticleResolver.CurArticles, middleware.BasicAuth(),
		middleware.LoginNeed(), schemabuilder.RelayConnection)
	// 获取文章（单个详细）
	query.FieldFunc("Article", resolve.ArticleResolver.Article)

	mutation := schema.Mutation()
	// 草稿
	mutation.FieldFunc("DraftArticle", resolve.ArticleResolver.Draft, middleware.BasicAuth(), middleware.LoginNeed())
	// 发布
	mutation.FieldFunc("NewArticle", resolve.ArticleResolver.NewArticle, middleware.BasicAuth(), middleware.LoginNeed())
	// 删除
	mutation.FieldFunc("DeleteArticle", resolve.ArticleResolver.Delete, middleware.BasicAuth(), middleware.LoginNeed())
}
