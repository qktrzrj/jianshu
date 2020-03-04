package controller

import "github.com/graphql-go/graphql"

var objtype = graphql.NewEnum(graphql.EnumConfig{
	Name: "Objtype",
	Values: graphql.EnumValueConfigMap{
		"article": {Value: "article", Description: "文章"},
		"comment": {Value: "comment", Description: "评论"},
		"reply":   {Value: "reply", Description: "回复"},
	},
	Description: "对象类型",
})
var zanType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Zan",
	Fields: graphql.Fields{
		"id":      {Type: graphql.NewNonNull(graphql.ID), Description: "ID"},
		"uid":     {Type: graphql.NewNonNull(graphql.ID), Description: "点赞用户ID"},
		"objtype": {Type: graphql.NewNonNull(objtype), Description: "被点赞对象类型"},
		"objid":   {Type: graphql.NewNonNull(graphql.ID), Description: "被点赞对象ID"},
	},
	Description: "点赞",
})
var zanListField = graphql.Field{
	Name: "zanList",
	Type: graphql.NewList(zanType),
	Args: graphql.FieldConfigArgument{
		"objtype": {Type: graphql.NewNonNull(objtype), Description: "被点赞对象类型"},
		"objid":   {Type: graphql.NewNonNull(graphql.ID), Description: "被点赞对象ID"},
	},
	Resolve:     nil,
	Description: "赞列表",
}

func registerZanType() {
	mutationType.AddFieldConfig("Zan", &graphql.Field{
		Type: zanType,
		Args: graphql.FieldConfigArgument{
			"objtype": {Type: graphql.NewNonNull(objtype), Description: "被点赞对象类型"},
			"objid":   {Type: graphql.NewNonNull(graphql.ID), Description: "被点赞对象ID"},
		},
		Resolve:     nil,
		Description: "点赞",
	})
	mutationType.AddFieldConfig("CancelZan", &graphql.Field{
		Type:        graphql.Boolean,
		Args:        graphql.FieldConfigArgument{"id": {Type: graphql.NewNonNull(graphql.ID), Description: "ID"}},
		Resolve:     nil,
		Description: "点赞",
	})
}
