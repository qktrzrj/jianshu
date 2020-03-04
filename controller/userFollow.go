package controller

import "github.com/graphql-go/graphql"

var userFollowType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserFollow",
	Fields: graphql.Fields{
		"id":   {Type: graphql.NewNonNull(graphql.ID), Description: "ID"},
		"uid":  {Type: graphql.NewNonNull(graphql.ID), Description: "用户ID"},
		"fuid": {Type: graphql.NewNonNull(graphql.ID), Description: "粉丝ID"},
	},
	Description: "用户关注关系",
})

func registerUserFollowType() {
	userType.AddFieldConfig("follows", &graphql.Field{
		Type:        graphql.NewList(userType),
		Resolve:     nil,
		Description: "用户关注列表",
	})
	userType.AddFieldConfig("fans", &graphql.Field{
		Type:        graphql.NewList(userType),
		Resolve:     nil,
		Description: "用户粉丝列表",
	})
	mutationType.AddFieldConfig("Follow", &graphql.Field{
		Type:        userFollowType,
		Args:        graphql.FieldConfigArgument{"uid": {Type: graphql.NewNonNull(graphql.ID), Description: "用户ID"}},
		Resolve:     nil,
		Description: "关注",
	})
	mutationType.AddFieldConfig("CancelFollow", &graphql.Field{
		Type:        userFollowType,
		Args:        graphql.FieldConfigArgument{"uid": {Type: graphql.NewNonNull(graphql.ID), Description: "用户ID"}},
		Resolve:     nil,
		Description: "取消关注",
	})
}
