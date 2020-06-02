package handler

import (
	"context"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/middleware"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/resolve"
)

func registerUser(schema *schemabuilder.Schema) {
	// 枚举类型映射
	schema.Enum("Gender", model.Gender(0), map[string]model.Gender{
		"Man":     model.Man,
		"Woman":   model.Woman,
		"Unknown": model.Unknown,
	})
	schema.Enum("UserState", model.UserState(0), map[string]model.UserState{
		"Unsigned":  model.Unsigned,
		"Forbidden": model.Forbidden,
		"Freeze":    model.Freeze,
	})

	// 将user结构体映射到graphql
	user := schema.Object("User", model.User{})
	// 粉丝数，关注数，文章数，字数，被点赞数
	user.FieldFunc("FansNum", func(u model.User) int { return u.Count.FansNum })
	user.FieldFunc("FollowNum", func(u model.User) int { return u.Count.FollowNum })
	user.FieldFunc("ArticleNum", func(u model.User) int { return u.Count.ArticleNum })
	user.FieldFunc("Words", func(u model.User) int { return u.Count.Words })
	user.FieldFunc("LikeNum", func(u model.User) int { return u.Count.LikeNum })

	query := schema.Query()
	// 获取用户信息
	query.FieldFunc("User", resolve.UserResolver.User)
	// 获取当前用户信息
	query.FieldFunc("CurrentUser", func(ctx context.Context) (model.User, error) {
		return resolve.UserResolver.User(ctx, resolve.IdArgs{Id: ctx.Value("userId").(int)})
	}, middleware.BasicAuth(), middleware.LoginNeed())
	// 校验用户名/邮箱唯一性
	query.FieldFunc("ValidUsername", resolve.UserResolver.ValidUsername)
	query.FieldFunc("ValidEmail", resolve.UserResolver.ValidEmail)
	// 粉丝列表
	query.FieldFunc("Fans", resolve.UserResolver.Followers)
	// 关注列表
	query.FieldFunc("Followed", resolve.UserResolver.Follows)
	// 用户关系
	query.FieldFunc("IsFollow", resolve.UserResolver.IsFollow, middleware.BasicAuth(), middleware.LoginNeed())

	mutation := schema.Mutation()
	// 注册
	mutation.FieldFunc("SignUp", resolve.UserResolver.SingUp, middleware.BasicAuth(), middleware.NotLogin())
	// 登录
	mutation.FieldFunc("SignIn", resolve.UserResolver.SignIn, middleware.BasicAuth(), middleware.NotLogin())
	// 退出登录
	mutation.FieldFunc("Logout", resolve.UserResolver.Logout, middleware.BasicAuth(), middleware.LoginNeed())
	// 关注
	mutation.FieldFunc("Follow", resolve.UserResolver.Follow, middleware.BasicAuth(), middleware.LoginNeed())
	// 取消关注
	mutation.FieldFunc("UnFollow", resolve.UserResolver.CancelFollow, middleware.BasicAuth(), middleware.LoginNeed())
	// 修改用户信息
	mutation.FieldFunc("UpdateUserInfo", resolve.UserResolver.UpdateUserInfo, middleware.BasicAuth(), middleware.LoginNeed())

}
