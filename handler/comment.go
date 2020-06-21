package handler

import (
	"context"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/resolve"
)

func registerComment(schema *schemabuilder.Schema) {
	comment := schema.Object("Comment", model.Comment{})

	// 评论人
	comment.FieldFunc("User", func(ctx context.Context, c model.Comment) (model.User, error) {
		return resolve.UserResolver.User(ctx, resolve.IdArgs{Id: c.Uid})
	})

	schema.Mutation().FieldFunc("AddComment", resolve.CommentResolver.Add, middleware.LoginNeed())
}
