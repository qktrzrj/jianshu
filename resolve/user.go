package resolve

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/jianshu/util"
	"github.com/shyptr/plugins/sqlog"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type userResolver struct{}

var UserResolver userResolver

type IdArgs struct {
	Id uint64 `graphql:"id"`
}

// 根据用户ID查询用户信息
func (u userResolver) User(ctx context.Context, args IdArgs) (model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	user, err := model.GetUser(tx, args.Id, "", "")
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, fmt.Errorf("查询用户信息失败")
	}
	count, err := model.GetUserCount(tx, args.Id)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, fmt.Errorf("查询用户信息失败")
	}
	user.Count = count
	return user, nil
}

// 粉丝列表
func (u userResolver) Followers(ctx context.Context, user model.User) ([]model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	ids, err := model.GetUserFollower(tx, user.Id)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, fmt.Errorf("查询用户信息失败")
	}
	var users []model.User
	for _, id := range ids {
		user, err := u.User(ctx, IdArgs{id})
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// 关注列表
func (u userResolver) Follows(ctx context.Context, user model.User) ([]model.User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	ids, err := model.GetFollowUser(tx, user.Id)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return nil, fmt.Errorf("查询用户信息失败")
	}
	var users []model.User
	for _, id := range ids {
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
		From(`"user"`).
		Where("username=$1", args.Username).
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
		From(`"user"`).
		Where("email=$1", args.Email).
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
func (u userResolver) SingUp(ctx context.Context, args model.UserArg) (user model.User, err error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

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

	args.Password = string(password)
	args.Avatar = "默认头像"
	id, err := model.InsertUser(tx, args)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, errors.New("注册失败")
	}
	err = model.InsertUserCount(tx, id)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, errors.New("注册失败")
	}

	// TODO:邮箱验证

	user, _ = model.GetUser(tx, id, "", "")
	return user, nil
}

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

	// 生成token，并设置cookie
	age := 1
	if args.RememberMe {
		age = 7 * 24
	}
	token, err := util.GeneraToken(user.Id, age)
	if err != nil {
		logger.Error().Caller().AnErr("生成token失败", err).Send()
		return model.User{}, errors.New("登录失败")
	}
	c := ctx.(*graphql.Context)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "me",
		Value:   token,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, 1),
		Domain:  "localhost",
		MaxAge:  int(time.Hour) * age * 24,
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
func (u userResolver) Follow(ctx context.Context, args struct {
	Id uint64 `graphql:"id"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(uint64)
	err := model.InsertUserFollow(tx, args.Id, userId)
	if err != nil {
		logger.Error().Caller().AnErr("关注失败", err).Send()
		return errors.New("关注失败")
	}
	// TODO: 发送通知
	return nil
}

// 取消关注
func (u userResolver) CancelFollow(ctx context.Context, args struct {
	Id uint64 `graphql:"id"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	userId := ctx.Value("userId").(uint64)
	err := model.DeleteUserFollow(tx, args.Id, userId)
	if err != nil {
		logger.Error().Caller().AnErr("取消关注失败", err).Send()
		return errors.New("取消关注失败")
	}

	return nil
}
