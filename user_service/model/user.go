package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type Gender int

const (
	Man Gender = iota + 1
	Woman
	Unknown
)

var GenderString = map[Gender]string{
	Man:     "man",
	Woman:   "woman",
	Unknown: "unknown",
}

type UserState int

const (
	Unsign UserState = iota + 1
	Normal
	Forbidden
	Freeze
)

var UserStateString = map[UserState]string{
	Unsign:    "unsign",
	Normal:    "normal",
	Forbidden: "forbidden",
	Freeze:    "freeze",
}

type User struct {
	Id        int64      `graphql:"id" db:"id"`
	Username  string     `graphql:"username" db:"username"`
	Email     string     `graphql:"email" db:"email"`
	Password  string     `graphql:"password" db:"password"`
	Avatar    string     `graphql:"avatar" db:"avatar"`
	Gender    Gender     `graphql:"Gender" db:"Gender"`
	Introduce *string    `graphql:"introduce" db:"introduce"`
	State     UserState  `graphql:"state" db:"state"`
	Root      bool       `graphql:"root" db:"root"`
	CreatedAt time.Time  `graphql:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `graphql:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `graphql:"deletedAt" db:"deleted_at"`
}

func GetUsers(ctx context.Context, where ...sqlex.Sqlex) ([]User, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	var users []User
	result, err := PSql.Select("id,username,email,password,avatar,gender,introduce,state,root,created_at,deleted_at").
		From(`"user"'`).
		WhereExpr(where...).RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, fmt.Errorf("fetch user list failed")
	}
	defer result.Close()
	for result.Next() {
		var user User
		err = result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.Gender, &user.Introduce, &user.State,
			&user.Root, &user.CreatedAt, &user.DeletedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, fmt.Errorf("fetch user list failed")
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(ctx context.Context, where ...sqlex.Sqlex) (User, error) {
	user := User{}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Select("id,username,email,password,avatar,gender,introduce,state,root,created_at,deleted_at").
		From(`"user"`).
		WhereExpr(where...).
		Limit(1).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return user, fmt.Errorf("fetch user info failed")
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.Gender, &user.Introduce, &user.State,
			&user.Root, &user.CreatedAt, &user.DeletedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return user, fmt.Errorf("fetch user info failed")
		}
	}
	return user, nil
}

func InsertUser(ctx context.Context, cv map[string]interface{}) (User, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return User{}, err
	}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	cv["id"] = int64(id)
	result, err := PSql.Insert(`"user"`).SetMap(cv).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return User{}, fmt.Errorf("save user info failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return User{}, fmt.Errorf("save user info failed")
	}
	return GetUser(ctx, sqlex.Eq{"id": id})
}

func UpdateUser(ctx context.Context, cv map[string]interface{}, where ...sqlex.Sqlex) error {
	for c, v := range cv {
		if v == nil {
			delete(cv, c)
		}
	}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	cv["updated_at"] = time.Now()
	result, err := PSql.Update(`"user"`).SetMap(cv).WhereExpr(where...).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("update user info failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("update user info failed")
	}
	return nil
}
