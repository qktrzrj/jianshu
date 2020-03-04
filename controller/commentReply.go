package controller

import "github.com/graphql-go/graphql"

var commentReplyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CommentReply",
	Fields: graphql.Fields{
		"id":      {Type: graphql.NewNonNull(graphql.ID), Description: "ID"},
		"cid":     {Type: graphql.NewNonNull(graphql.ID), Description: "评论ID"},
		"uid":     {Type: graphql.NewNonNull(graphql.ID), Description: "回复人ID"},
		"content": {Type: graphql.NewNonNull(graphql.String), Description: "内容"},
		"state":   {Type: graphql.NewNonNull(articleState), Description: "状态"},
	},
	Description: "评论回复",
})

func registerCommentReplyType() {
	mutationType.AddFieldConfig("Reply", &graphql.Field{
		Type: commentReplyType,
		Args: graphql.FieldConfigArgument{
			"cid":     {Type: graphql.NewNonNull(graphql.ID), Description: "评论ID"},
			"content": {Type: graphql.NewNonNull(graphql.String), Description: "内容"},
		},
		Resolve:     nil,
		Description: "回复",
	})
	mutationType.AddFieldConfig("DeleteReply", &graphql.Field{
		Type:        commentReplyType,
		Args:        graphql.FieldConfigArgument{"id": {Type: graphql.NewNonNull(graphql.ID), Description: "ID"}},
		Resolve:     nil,
		Description: "删除回复",
	})
}
