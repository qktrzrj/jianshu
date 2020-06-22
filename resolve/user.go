package resolve

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql"
	"github.com/shyptr/jianshu/cache"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/util"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type userResolver struct{}

var UserResolver userResolver

type IdArgs struct {
	Id int `graphql:"id"`
}

func (u userResolver) Users(ctx context.Context, arg struct {
	Username string `graphql:"username;;null"`
}) ([]model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	users, err := model.GetUsers(tx, arg.Username)

	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, fmt.Errorf("查询用户信息失败")
	}

	return users, nil
}

// 根据用户ID查询用户信息
func (u userResolver) User(ctx context.Context, args IdArgs) (model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	user, err := cache.QueryCache(ctx, cache.User{Id: args.Id}, func() (interface{}, error) {
		return model.GetUser(tx, args.Id, "", "")
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, fmt.Errorf("查询用户信息失败")
	}
	count, err := cache.QueryCache(ctx, cache.UserCount{Uid: args.Id}, func() (interface{}, error) {
		return model.GetUserCount(tx, args.Id)
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, fmt.Errorf("查询用户信息失败")
	}
	res := user.(model.User)
	res.Count = count.(model.UserCount)
	return res, nil
}

// 粉丝列表
func (u userResolver) Followers(ctx context.Context, arg IdArgs) ([]model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	ids, err := cache.QueryCaches(ctx, cache.Follow{Uid: arg.Id}, func() (interface{}, error) {
		return model.GetUserFollower(tx, arg.Id)
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, fmt.Errorf("获取粉丝列表失败")
	}
	var users []model.User
	for _, id := range ids.([]int) {
		user, err := u.User(ctx, IdArgs{id})
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// 关注列表
func (u userResolver) Follows(ctx context.Context, arg IdArgs) ([]model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	ids, err := cache.QueryCaches(ctx, cache.Follow{Fuid: arg.Id}, func() (interface{}, error) {
		return model.GetFollowUser(tx, arg.Id)
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, fmt.Errorf("获取粉丝列表失败")
	}
	var users []model.User
	for _, id := range ids.([]int) {
		user, err := u.User(ctx, IdArgs{id})
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

type usernameArg struct {
	Username string `graphql:"username"`
}

// 校验用户名唯一性
func (u userResolver) ValidUsername(ctx context.Context, args usernameArg) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	rows, err := model.PSql.
		Select("count(id)").
		From("`user`").
		Where(sqlex.Eq{"username": args.Username}).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return errors.New("校验用户名失败")
	}
	var count int
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return errors.New("校验用户名失败")
		}
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}
	return nil
}

type emailArg struct {
	Email string `graphql:"email" validate:"email"`
}

// 校验邮箱唯一性
func (u userResolver) ValidEmail(ctx context.Context, args emailArg) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	rows, err := model.PSql.
		Select("count(id)").
		From("`user`").
		Where(sqlex.Eq{"email": args.Email}).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return errors.New("校验邮箱失败")
	}
	var count int
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			return errors.New("校验邮箱失败")
		}
	}
	if count > 0 {
		return errors.New("邮箱已被使用")
	}
	return nil
}

// 注册
func (u userResolver) SingUp(ctx context.Context, args struct {
	Username string `graphql:"username" validate:"min=6,max=16"`
	Email    string `graphql:"email" validate:"email"`
	Password string `graphql:"password" validate:"min=8"`
}) (user model.User, err error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	// 校验用户名和邮箱信息
	err = u.ValidUsername(ctx, usernameArg{Username: args.Username})
	if err != nil {
		return model.User{}, err
	}
	err = u.ValidEmail(ctx, emailArg{Email: args.Email})
	if err != nil {
		return model.User{}, err
	}

	// 密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(args.Password), 10)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, errors.New("注册失败")
	}

	id, err := model.InsertUser(tx, map[string]interface{}{
		"username": args.Username,
		"email":    args.Email,
		"password": string(password),
		"avatar":   "http://www.shyptr.cn/image/default",
		"root":     false,
	})
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, errors.New("注册失败")
	}

	// TODO:邮箱验证

	return u.User(ctx, IdArgs{id})
}

// 登陆
func (u userResolver) SignIn(ctx context.Context, args struct {
	Username   string `graphql:"username"` // 邮箱或者用户名
	Password   string `graphql:"password"`
	RememberMe bool   `graphql:"rememberme"`
}) (user model.User, err error) {
	// 验证账号密码
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	user, err = model.GetUser(tx, 0, args.Username, args.Username)
	if err != nil {
		logger.Error().Caller().AnErr("登录失败", err).Send()
		return model.User{}, errors.New("登录失败")
	}

	if user.Id == 0 {
		return model.User{}, errors.New("用户不存在！")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(args.Password))
	if err != nil {
		return model.User{}, err
	}

	user.Count, _ = model.GetUserCount(tx, user.Id)

	// 生成token，并设置cookie,默认一天,记住则为7天
	// TODO: remember天数应该做配置化
	age := time.Hour * 24
	if args.RememberMe {
		age *= 7
	}
	token, err := util.GeneraToken(user.Id, user.Root, int(user.State), age)
	if err != nil {
		logger.Error().Caller().AnErr("生成token失败", err).Send()
		return model.User{}, errors.New("登录失败")
	}
	c := ctx.(*graphql.Context)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "me",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(age),
		Domain:  "localhost",
	})
	return user, nil
}

func (u userResolver) Logout(ctx context.Context) {
	c := ctx.(*graphql.Context)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "me",
		Path:   "/",
		Domain: "localhost",
		MaxAge: -1,
	})
}

// 关注
func (u userResolver) Follow(ctx context.Context, args IdArgs) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)
	// 删除缓存
	cache.Delete(cache.Follow{Uid: args.Id}.GetCachesKey())
	cache.Delete(cache.Follow{Fuid: userId}.GetCachesKey())
	cache.Delete(cache.UserCount{Uid: userId}.GetCacheKey())
	cache.Delete(cache.UserCount{Uid: args.Id}.GetCacheKey())

	err := model.InsertUserFollow(tx, args.Id, userId)
	if err != nil {
		logger.Error().Caller().AnErr("关注失败", err).Send()
		return errors.New("关注失败")
	}
	err = model.UpdateUserCount(tx, userId, 1, true)
	if err != nil {
		logger.Error().Caller().AnErr("关注失败", err).Send()
		return errors.New("关注失败")
	}
	err = model.UpdateUserCount(tx, args.Id, 0, true)
	if err != nil {
		logger.Error().Caller().AnErr("关注失败", err).Send()
		return errors.New("关注失败")
	}
	// TODO: 发送通知
	return nil
}

// 取消关注
func (u userResolver) CancelFollow(ctx context.Context, args IdArgs) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	// 删除缓存
	cache.Delete(cache.Follow{Uid: args.Id}.GetCachesKey())
	cache.Delete(cache.Follow{Fuid: userId}.GetCachesKey())
	cache.Delete(cache.UserCount{Uid: userId}.GetCacheKey())
	cache.Delete(cache.UserCount{Uid: args.Id}.GetCacheKey())

	err := model.DeleteUserFollow(tx, args.Id, userId)
	if err != nil {
		logger.Error().Caller().AnErr("取消关注失败", err).Send()
		return errors.New("取消关注失败")
	}
	err = model.UpdateUserCount(tx, userId, 1, false)
	if err != nil {
		logger.Error().Caller().AnErr("取消关注失败", err).Send()
		return errors.New("取消关注失败")
	}
	err = model.UpdateUserCount(tx, args.Id, 0, false)
	if err != nil {
		logger.Error().Caller().AnErr("取消关注失败", err).Send()
		return errors.New("取消关注失败")
	}

	return nil
}

// 修改用户信息
func (u userResolver) UpdateUserInfo(ctx context.Context, arg struct {
	Username  *string       `graphql:"username"`
	Email     *string       `graphql:"email"`
	Password  *string       `graphql:"password"`
	Avatar    *string       `graphql:"avatar"`
	Gender    *model.Gender `graphql:"gender"`
	Introduce *string       `graphql:"introduce"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)

	// 删除缓存
	cache.Delete(cache.User{Id: userId}.GetCacheKey())

	setMap := make(map[string]interface{})
	if arg.Email != nil {
		setMap["email"] = *arg.Email
		if *arg.Email == "" {
			return errors.New("邮箱不能为空！")
		}
	}
	if arg.Password != nil {
		setMap["password"] = *arg.Password
	}
	if arg.Avatar != nil {
		setMap["avatar"] = *arg.Avatar
	}
	if arg.Gender != nil {
		setMap["gender"] = *arg.Gender
	}
	if arg.Introduce != nil {
		setMap["introduce"] = *arg.Introduce
	}
	if arg.Username != nil {
		if *arg.Username == "" {
			return errors.New("用户名不能为空！")
		}
		setMap["username"] = *arg.Username
	}

	err := model.UpdateUser(tx, userId, setMap)
	if err != nil {
		logger.Error().Caller().AnErr("修改用户信息失败", err).Send()
		return errors.New("修改用户信息失败")
	}

	return nil
}

// 用户关系
func (u userResolver) IsFollow(ctx context.Context, arg IdArgs) (bool, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(int)
	isFollow, err := model.IsFollow(tx, arg.Id, userId)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return false, errors.New("查询用户关系失败")
	}
	return isFollow, nil
}
