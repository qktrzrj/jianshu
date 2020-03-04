package controller

import "github.com/graphql-go/graphql"

var tagType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tag",
	Fields: graphql.Fields{
		"id":   {Type: graphql.NewNonNull(graphql.ID), Description: "ID"},
		"name": {Type: graphql.NewNonNull(graphql.String), Description: "标签名"},
	},
	Description: "标签",
})

func registerTagType() {
	queryType.AddFieldConfig("Tag", &graphql.Field{
		Type:        graphql.NewList(tagType),
		Args:        graphql.FieldConfigArgument{"name": {Type: graphql.String, Description: "标签名"}},
		Resolve:     nil,
		Description: "获取标签列表",
	})
	mutationType.AddFieldConfig("AddTag", &graphql.Field{
		Type:        tagType,
		Args:        graphql.FieldConfigArgument{"name": {Type: graphql.NewNonNull(graphql.String), Description: "标签名"}},
		Resolve:     nil,
		Description: "添加新标签",
	})
}
