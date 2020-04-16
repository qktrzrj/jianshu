package resolve

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/hello-world-web/user_service/model"
	"github.com/shyptr/hello-world-web/util"
	"github.com/shyptr/sqlex"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	metadata2 "google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

var validate = schemabuilder.NewValidate()

type userResolve struct{}

var UserResolve = userResolve{}

func (this userResolve) GetUserList(ctx context.Context, args struct {
	Username *string `graphql:"username"`
	Email    *string `graphql:"email"`
}) ([]model.User, error) {
	if args.Email != nil {
		if err := validate.Var(*args.Email, "email"); err != nil {
			return nil, err
		}
	}
	return model.GetUsers(ctx, []sqlex.Sqlex{
		sqlex.IF{Condition: args.Username != nil, Sq: sqlex.Like{"username": args.Username}},
		sqlex.IF{Condition: args.Email != nil, Sq: sqlex.Like{"email": args.Email}},
	}...)
}

func (this userResolve) GetUser(ctx context.Context, args struct {
	Id       *int64  `graphql:"id"`
	Username *string `graphql:"username"`
}) (model.User, error) {
	return model.GetUser(ctx, []sqlex.Sqlex{
		sqlex.IF{Condition: args.Id != nil, Sq: sqlex.Eq{"id": args.Id}},
		sqlex.IF{Condition: args.Username != nil, Sq: sqlex.Like{"username": args.Username}},
	}...)
}

func (this userResolve) CurrentUser(ctx context.Context) (model.User, error) {
	id := ctx.Value("userId")
	return model.GetUser(ctx, sqlex.Eq{"id": id})
}

func (this userResolve) GetUserCount(ctx context.Context, source model.User) (model.UserCount, error) {
	return model.GetUserCount(ctx, source.Id)
}

func (this userResolve) CheckUsername(ctx context.Context, args struct {
	Username string `graphql:"username" validate:"min=6,max=20"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	rows, err := model.PSql.
		Select("count(id)").
		From(`"user"`).
		Where(sqlex.Eq{"username": args.Username}).
		RunWith(model.DB).Query()

	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("validate username failed")
	}
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			logger.Error().Err(err).Send()
			return errors.New("validate username failed")
		}
	}
	if count > 0 {
		return errors.New("username exist")
	}
	return nil
}

func (this userResolve) CheckEmail(ctx context.Context, args struct {
	Email string `graphql:"email" validate:"email"`
}) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	rows, err := model.PSql.
		Select("count(id) as c").
		From(`"user"`).
		Where(sqlex.Eq{"email": args.Email}).
		RunWith(model.DB).Query()

	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("validate email failed")
	}
	var count int
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			logger.Error().Err(err).Send()
			return errors.New("validate email failed")
		}
	}
	if count > 0 {
		return errors.New("email exist")
	}
	return nil
}

func (this userResolve) SignUp(ctx context.Context, args struct {
	Username string `graphql:"username" validate:"max=20,min=6"`
	Email    string `graphql:"email" validate:"email"`
	Password string `graphql:"password" validate:"max=20,min=8"`
}) (model.User, error) {
	err := this.CheckUsername(ctx, struct {
		Username string `graphql:"username" validate:"min=6,max=20"`
	}(struct{ Username string }{Username: args.Username}))
	if err != nil {
		return model.User{}, err
	}
	err = this.CheckEmail(ctx, struct {
		Email string `graphql:"email" validate:"email"`
	}(struct{ Email string }{Email: args.Email}))
	if err != nil {
		return model.User{}, err
	}
	logger := ctx.Value("logger").(zerolog.Logger)
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), 10)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return model.User{}, errors.New("sign up failed")
	}
	user, err := model.InsertUser(ctx, map[string]interface{}{
		"username": args.Username,
		"email":    args.Email,
		"password": generateFromPassword,
		"avatar":   "",
	})
	if err != nil {
		return model.User{}, err
	}
	err = model.InsertUserCount(ctx, user.Id)
	if err != nil {
		return model.User{}, err
	}
	//TODO: send sign email

	return user, nil
}

func (this userResolve) SignIn(ctx context.Context, args struct {
	Username *string `graphql:"username"`
	Email    *string `graphql:"email"`
	Password string  `graphql:"password"`
	Remember bool    `graphql:"remember"`
}) (model.User, error) {
	if args.Username == nil && args.Email == nil {
		return model.User{}, fmt.Errorf("username or email must provide one")
	}
	user, err := model.GetUser(ctx, []sqlex.Sqlex{
		sqlex.IF{Condition: args.Username != nil, Sq: sqlex.Eq{"username": args.Username}},
		sqlex.IF{Condition: args.Email != nil, Sq: sqlex.Eq{"email": args.Email}},
	}...)
	if err != nil {
		return model.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(args.Password))
	if err != nil {
		return model.User{}, err
	}
	var token string
	age := 1
	if args.Remember {
		age = 7
	}
	token, err = util.GenToken(user.Id, age)
	if err != nil {
		return model.User{}, err
	}

	err = grpc.SendHeader(ctx, metadata2.New(map[string]string{
		"cookie": (&http.Cookie{
			Name:    "me",
			Value:   token,
			Path:    "/",
			Expires: time.Now().Add(time.Hour * time.Duration(24*age)),
			MaxAge:  int(time.Now().Add(time.Hour * time.Duration(24*age)).Unix()),
		}).String(),
	}))
	return user, err
}

func (this userResolve) Logout(ctx context.Context) {
	ctx = metadata.Set(ctx, "cookie", (&http.Cookie{
		Name:    "me",
		Path:    "/",
		Expires: time.Now().Add(time.Second * 0),
		MaxAge:  0,
	}).String())
}

func (this userResolve) Update(ctx context.Context, args struct {
	Id        int64         `graphql:"id"`
	Username  *string       `graphql:"username"`
	Email     *string       `graphql:"email"`
	Avatar    *string       `graphql:"avatar"`
	Gender    *model.Gender `graphql:"gender"`
	Introduce *string       `graphql:"introduce"`
}) (model.User, error) {
	return model.User{}, model.UpdateUser(ctx, map[string]interface{}{
		"username":  args.Username,
		"email":     args.Email,
		"avatar":    args.Avatar,
		"gender":    args.Gender,
		"introduce": args.Introduce,
	}, sqlex.Eq{"id": args.Id})
}

func (this userResolve) Follow(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) error {
	return model.InsertUserFollow(ctx, args.Id, ctx.Value("userId").(int64))
}

func (this userResolve) UnFollow(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) error {
	return model.RemoveUserFollow(ctx, args.Id, ctx.Value("userId").(int64))
}

func (this userResolve) FollowList(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) ([]model.User, error) {
	followList, err := model.GetUserFollowList(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for _, uid := range followList {
		user, err := model.GetUser(ctx, sqlex.Eq{"id": uid})
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (this userResolve) FollowerList(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) ([]model.User, error) {
	followList, err := model.GetFollowUserList(ctx, args.Id)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for _, uid := range followList {
		user, err := model.GetUser(ctx, sqlex.Eq{"id": uid})
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
