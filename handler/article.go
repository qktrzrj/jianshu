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
	schema.Enum("ArticleState", model.ArticleState(0), map[string]schemabuilder.DescField{
		"Draft":     {model.Draft, "草稿"},
		"Unaudited": {model.Unaudited, "未审核"},
		"Online":    {model.Online, "已发布"},
		"Offline":   {model.Offline, "已下线"},
		"Deleted":   {model.Deleted, "已删除"},
		"Updated":   {model.Updated, "更新未重新发布"},
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
	// 文章评论
	article.FieldFunc("CommentList", resolve.CommentResolver.List)

	query := schema.Query()
	schemabuilder.RelayKey(model.Article{}, "id")
	schema.Object("articlePage", resolve.ArticlePage{})
	// 热门文章（分页）
	query.FieldFunc("HotArticles", resolve.ArticleResolver.Hots, schemabuilder.RelayConnection)
	// 查询文章（分页）
	query.FieldFunc("Articles", resolve.ArticleResolver.Articles, schemabuilder.RelayConnection)
	// 获取登录人文章（分页）
	query.FieldFunc("CurArticles", resolve.ArticleResolver.CurArticles,
		middleware.LoginNeed(), schemabuilder.RelayConnection)
	// 获取登录人喜欢的文章(分页)
	query.FieldFunc("CurLikeArticles", resolve.ArticleResolver.LikeArticles, middleware.LoginNeed(),
		schemabuilder.RelayConnection)
	// 获取文章（单个详细）
	query.FieldFunc("Article", resolve.ArticleResolver.Article)

	mutation := schema.Mutation()
	// 草稿
	mutation.FieldFunc("DraftArticle", resolve.ArticleResolver.Draft, middleware.LoginNeed())
	// 发布
	mutation.FieldFunc("NewArticle", resolve.ArticleResolver.NewArticle, middleware.LoginNeed())
	// 修改
	mutation.FieldFunc("UpdateArticle", resolve.ArticleResolver.UpdateArticle, middleware.LoginNeed())
	// 删除
	mutation.FieldFunc("DeleteArticle", resolve.ArticleResolver.Delete, middleware.LoginNeed())
	// 文章浏览数
	mutation.FieldFunc("ViewAdd", resolve.ArticleResolver.View)
}
