package model

import (
	"context"
	"errors"
	"github.com/unrotten/builder"
	"time"
)

const (
	Man     = "man"
	Woman   = "woman"
	Unknown = "unknown"
)

const (
	Unsign    = "unsign"
	Normal    = "normal"
	Forbidden = "forbidden"
	Freeze    = "freeze"
)

type User struct {
	Id        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Gender    string    `json:"gender" db:"gender"`
	Introduce string    `json:"introduce" db:"introduce"`
	State     string    `json:"state" db:"state"`
	Root      bool      `json:"root" db:"root"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetUsers(ctx context.Context, where where) ([]User, error) {
	result := selectList(ctx, `"user"`, where)
	if !result.success {
		return nil, errors.New("获取用户列表失败")
	}
	list, ok := builder.Get(result.b, "list")
	if !ok {
		return nil, errors.New("获取用户列表失败")
	}
	users := make([]User, 0, len(list.([]interface{})))
	for _, item := range list.([]interface{}) {
		users = append(users, builder.GetStructLikeByTag(item.(builder.Builder), User{}, "db").(User))
	}
	return users, nil
}

func GetUser(ctx context.Context, where where) (User, error) {
	result := selectOne(ctx, `"user"`, where)
	if !result.success {
		return User{}, errors.New("查询用户数据失败")
	}
	return builder.GetStructLikeByTag(result.b, User{}, "db").(User), nil
}

func InsertUser(ctx context.Context, cv map[string]interface{}) (User, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return User{}, err
	}

	cv["id"] = int64(id)
	result := insertOne(ctx, `"user"`, cv)
	if !result.success {
		return User{}, errors.New("插入用户数据失败")
	}
	return builder.GetStructLikeByTag(result.b, User{}, "db").(User), nil
}

func UpdateUser(ctx context.Context, cv cv, where where) error {
	result := update(ctx, `"user"`, cv, where)
	if !result.success {
		return errors.New("更新用户数据失败")
	}
	return nil
}
