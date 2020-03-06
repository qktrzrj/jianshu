package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/unrotten/hello-world-web/model"
)

var gender = graphql.NewEnum(graphql.EnumConfig{
	Name: "Gender",
	Values: graphql.EnumValueConfigMap{
		"man":     {Value: model.Man, Description: "男"},
		"woman":   {Value: model.Woman, Description: "女"},
		"unknown": {Value: model.Unknown, Description: "保密"},
	},
	Description: "性别",
})
var userState = graphql.NewEnum(graphql.EnumConfig{
	Name: "UserState",
	Values: graphql.EnumValueConfigMap{
		"unsign":    {Value: model.Unsign, Description: "未认证"},
		"normal":    {Value: model.Normal, Description: "正常"},
		"forbidden": {Value: model.Forbidden, Description: "禁止发言"},
		"freeze":    {Value: model.Freeze, Description: "冻结"},
	},
	Description: "用户状态",
})

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        relay.GlobalIDField("User", nil),
		"username":  {Type: graphql.NewNonNull(graphql.String), Description: "用户名"},
		"email":     {Type: graphql.NewNonNull(graphql.String), Description: "邮箱"},
		"avatar":    {Type: graphql.NewNonNull(graphql.String), Description: "头像"},
		"gender":    {Type: graphql.NewNonNull(gender), Description: "性别"},
		"introduce": {Type: graphql.String, Description: "个人简介"},
		"state":     {Type: graphql.NewNonNull(userState), Description: "状态"},
		"root":      {Type: graphql.NewNonNull(graphql.Boolean), Description: "管理员"},
		"userCount": {
			Type:        userCountType,
			Resolve:     nil,
			Description: "用户计数信息",
		},
	},
	Description: "用户数据",
})
var userConnectionDefinition = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "User",
	NodeType: userType,
})

func registerUserType() {
	queryType.AddFieldConfig("GetUserList", &graphql.Field{
		Type:        userConnectionDefinition.ConnectionType,
		Args:        relay.ConnectionArgs,
		Resolve:     nil,
		Description: "用户列表",
	})
	queryType.AddFieldConfig("GetUser", &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id":       {Type: graphql.ID, Description: "ID"},
			"username": {Type: graphql.String, Description: "用户名"},
		},
		Resolve:     nil,
		Description: "获取用户信息",
	})
	queryType.AddFieldConfig("CurrentUser", &graphql.Field{
		Type:        userType,
		Resolve:     nil,
		Description: "获取当前登录用户信息",
	})
	mutationType.AddFieldConfig("CreatUser", &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"username": {Type: graphql.NewNonNull(graphql.String), Description: "用户名"},
			"email":    {Type: graphql.NewNonNull(graphql.String), Description: "邮箱"},
			"password": {Type: graphql.NewNonNull(graphql.String), Description: "密码"},
		},
		Resolve:     nil,
		Description: "注册新用户",
	})
	mutationType.AddFieldConfig("SignIn", &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"username": {Type: graphql.NewNonNull(graphql.String), Description: "用户名"},
			"password": {Type: graphql.NewNonNull(graphql.String), Description: "密码"},
		},
		Resolve:     nil,
		Description: "用户登录",
	})
	mutationType.AddFieldConfig("UpdateUser", &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"username":  {Type: graphql.String, Description: "用户名"},
			"email":     {Type: graphql.String, Description: "邮箱"},
			"avatar":    {Type: graphql.String, Description: "头像"},
			"gender":    {Type: gender, Description: "性别"},
			"introduce": {Type: graphql.String, Description: "个人简介"},
		},
		Resolve:     nil,
		Description: "修改用户信息",
	})
}
