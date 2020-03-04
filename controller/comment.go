package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var commentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id":      relay.GlobalIDField("Comment", nil),
		"aid":     {Type: graphql.NewNonNull(graphql.String), Description: "文章ID"},
		"uid":     {Type: graphql.NewNonNull(graphql.String), Description: "评论用户ID"},
		"content": {Type: graphql.NewNonNull(graphql.String), Description: "评论内容"},
		"zanNum":  {Type: graphql.NewNonNull(graphql.Int), Description: "被赞数"},
		"floor":   {Type: graphql.NewNonNull(graphql.Int), Description: "第几楼"},
		"state":   {Type: graphql.NewNonNull(articleState), Description: "状态"},
		"replies": {
			Type:        graphql.NewList(commentReplyType),
			Resolve:     nil,
			Description: "评论回复列表",
		},
	},
	Description: "评论",
})
var commentConnectDefinitions = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "Comment",
	NodeType: commentType,
})

func registerCommentType() {
	queryType.AddFieldConfig("Comments", &graphql.Field{
		Type: commentConnectDefinitions.ConnectionType,
		Args: relay.NewConnectionArgs(graphql.FieldConfigArgument{
			"aid": {Type: graphql.NewNonNull(graphql.ID), Description: "文章ID"},
		}),
		Resolve:     nil,
		Description: "获取文章评论",
	})
	mutationType.AddFieldConfig("Comment", &graphql.Field{
		Type: commentType,
		Args: graphql.FieldConfigArgument{
			"aid":     {Type: graphql.NewNonNull(graphql.ID), Description: "文章ID"},
			"content": {Type: graphql.NewNonNull(graphql.String), Description: "评论内容"},
		},
		Resolve:     nil,
		Description: "评论",
	})
	mutationType.AddFieldConfig("DeleteComment", &graphql.Field{
		Type:        commentType,
		Args:        graphql.FieldConfigArgument{"id": {Type: graphql.NewNonNull(graphql.ID), Description: "ID"}},
		Resolve:     nil,
		Description: "删除评论",
	})
}
