package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var articleState = graphql.NewEnum(graphql.EnumConfig{
	Name: "ArticleState",
	Values: graphql.EnumValueConfigMap{
		"unaudited": {Value: "unaudited", Description: "未审核"},
		"online":    {Value: "online", Description: "已上线"},
		"offline":   {Value: "offline", Description: "已下线"},
		"deleted":   {Value: "deleted", Description: "已删除"},
	},
	Description: "文章状态",
})
var articleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id":      relay.GlobalIDField("Article", nil),
		"sn":      {Type: graphql.NewNonNull(graphql.String), Description: "序号"},
		"title":   {Type: graphql.NewNonNull(graphql.String), Description: "标题"},
		"uid":     {Type: graphql.NewNonNull(graphql.ID), Description: "作者ID"},
		"cover":   {Type: graphql.String, Description: "封面"},
		"content": {Type: graphql.String, Description: "文章内容"},
		"tags":    {Type: graphql.NewList(graphql.String), Description: "标签"},
		"state":   {Type: graphql.NewNonNull(articleState), Description: "状态"},
		"count": {
			Type:        articleExType,
			Resolve:     nil,
			Description: "文章计数",
		},
	},
	Description: "文章",
})
var articleConnectionDefinition = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "Article",
	NodeType: articleType,
})

func registerArticleType() {
	queryType.AddFieldConfig("GetArticle", &graphql.Field{
		Type:        articleType,
		Args:        graphql.FieldConfigArgument{"id": {Type: graphql.NewNonNull(graphql.ID), Description: "ID"}},
		Resolve:     nil,
		Description: "获取指定文章",
	})
	queryType.AddFieldConfig("Articles", &graphql.Field{
		Type: articleConnectionDefinition.ConnectionType,
		Args: relay.NewConnectionArgs(graphql.FieldConfigArgument{
			"title":   {Type: graphql.String, Description: "标题"},
			"uid":     {Type: graphql.ID, Description: "作者ID"},
			"content": {Type: graphql.String, Description: "内容"},
			"tags":    {Type: graphql.NewList(graphql.String), Description: "标签"},
		}),
		Resolve:     nil,
		Description: "获取文章列表",
	})
	mutationType.AddFieldConfig("CreateArticle", &graphql.Field{
		Type: articleType,
		Args: graphql.FieldConfigArgument{
			"title":   {Type: graphql.NewNonNull(graphql.String), Description: "标题"},
			"cover":   {Type: graphql.String, Description: "封面"},
			"content": {Type: graphql.String, Description: "文章内容"},
			"tags":    {Type: graphql.NewList(graphql.String), Description: "标签"},
		},
		Resolve:     nil,
		Description: "新增文章",
	})
	mutationType.AddFieldConfig("UpdateArticle", &graphql.Field{
		Type: articleType,
		Args: graphql.FieldConfigArgument{
			"id":      {Type: graphql.NewNonNull(graphql.ID), Description: "ID"},
			"title":   {Type: graphql.NewNonNull(graphql.String), Description: "标题"},
			"cover":   {Type: graphql.String, Description: "封面"},
			"content": {Type: graphql.String, Description: "文章内容"},
			"tags":    {Type: graphql.NewList(graphql.String), Description: "标签"},
		},
		Resolve:     nil,
		Description: "修改文章",
	})
	mutationType.AddFieldConfig("DeleteArticle", &graphql.Field{
		Type:        articleType,
		Args:        graphql.FieldConfigArgument{"id": {Type: graphql.NewNonNull(graphql.ID), Description: "ID"}},
		Resolve:     nil,
		Description: "删除文章",
	})
}
