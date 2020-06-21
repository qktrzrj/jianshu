package handler

import (
	"context"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/resolve"
)

func registerReply(schema *schemabuilder.Schema) {
	reply := schema.Object("Reply", model.Reply{})

	reply.FieldFunc("User", func(ctx context.Context, r model.Reply) (model.User, error) {
		return resolve.UserResolver.User(ctx, resolve.IdArgs{Id: r.Uid})
	})

	schema.Query().FieldFunc("ReplyList", resolve.ReplyResolver.List)

	schema.Mutation().FieldFunc("AddReply", resolve.ReplyResolver.Add, middleware.LoginNeed())
}
