package handler

import (
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/resolve"
)

func registerLike(schema *schemabuilder.Schema) {
	schema.Enum("ObjType", model.Objtype(0), map[string]model.Objtype{
		"ArticleObj": model.ArticleObj,
		"CommentObj": model.CommentObj,
		"ReplyObj":   model.ReplyObj,
	})

	schema.Query().FieldFunc("HasLike", resolve.LikeResolver.HasLike, middleware.LoginNeed())

	mutation := schema.Mutation()
	mutation.FieldFunc("Like", resolve.LikeResolver.Like, middleware.LoginNeed())
	mutation.FieldFunc("Unlike", resolve.LikeResolver.UnLike, middleware.LoginNeed())
}
