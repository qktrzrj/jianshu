package handler

import (
	"context"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/resolve"
)

func registerMsg(schema *schemabuilder.Schema) {
	schema.Enum("MsgType", model.MsgType(0), map[string]model.MsgType{
		"CommentMsg": model.CommentMsg,
		"ReplyMsg":   model.ReplyMsg,
		"LikeMsg":    model.LikeMsg,
		"FollowMsg":  model.FollowMsg,
	})

	schema.Object("MsgNum", model.MsgNum{})
	msg := schema.Object("Msg", model.Msg{})

	msg.FieldFunc("User", func(ctx context.Context, m model.Msg) (model.User, error) {
		return resolve.UserResolver.User(ctx, resolve.IdArgs{Id: m.FromId})
	})

	schema.Query().FieldFunc("MsgNum", resolve.MsgResolver.MsgNum, middleware.LoginNeed())
	schema.Query().FieldFunc("ListMsg", resolve.MsgResolver.ListMsg, middleware.LoginNeed())

	schema.Mutation().FieldFunc("AddMsg", resolve.MsgResolver.AddMsg, middleware.LoginNeed())
}
