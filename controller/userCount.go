package controller

import "github.com/graphql-go/graphql"

var userCountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserCount",
	Fields: graphql.Fields{
		"uid":        {Type: graphql.NewNonNull(graphql.ID), Description: "用户ID"},
		"fansNum":    {Type: graphql.NewNonNull(graphql.Int), Description: "用户粉丝数"},
		"followNum":  {Type: graphql.NewNonNull(graphql.Int), Description: "用户关注数(关注其他用户)"},
		"articleNum": {Type: graphql.NewNonNull(graphql.Int), Description: "用户文章数"},
		"words":      {Type: graphql.NewNonNull(graphql.Int), Description: "字数"},
		"zanNum":     {Type: graphql.NewNonNull(graphql.Int), Description: "用户被赞总数"},
	},
	Description: "用户相关计数信息",
})
