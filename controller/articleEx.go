package controller

import "github.com/graphql-go/graphql"

var articleExType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArticleEx",
	Fields: graphql.Fields{
		"aid":     {Type: graphql.NewNonNull(graphql.ID), Description: "文章ID"},
		"viewNum": {Type: graphql.NewNonNull(graphql.Int), Description: "浏览数"},
		"cmtNum":  {Type: graphql.NewNonNull(graphql.Int), Description: "评论数"},
		"zanNum":  {Type: graphql.NewNonNull(graphql.Int), Description: "点赞数"},
	},
	Description: "文章计数信息",
})
