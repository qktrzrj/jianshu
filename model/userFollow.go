package model

import (
	"context"
	"errors"
	"github.com/unrotten/sqlex"
	"time"
)

type UserFollow struct {
	Id        int64     `json:"id" db:"id"`
	Uid       int64     `json:"uid" db:"uid"`
	Fuid      int64     `json:"fuid" db:"fuid"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func InsertUserFollow(ctx context.Context, uid, fuid int64) error {
	id, err := idfetcher.NextID()
	if err != nil {
		return err
	}
	if result := insertOne(ctx, UserFollow{}, "user_follow", cv{"id": int64(id), "uid": uid, "fuid": fuid}); !result.success {
		return errors.New("插入用户关注表失败")
	}
	return nil
}

func RemoveUserFollow(ctx context.Context, uid, fuid int64) error {
	if !remove(ctx, "user_follow", where{sqlex.Eq{"uid": uid, "fuid": fuid}}).success {
		return errors.New("删除用户关注失败")
	}
	return nil
}

// 获取用户关注列表
func GetUserFollowList(ctx context.Context, fuid int64) ([]UserFollow, error) {
	result := selectList(ctx, int64(0), "user_follow", where{sqlex.Eq{"fuid": fuid}}, "uid")
	if !result.success {
		return nil, errors.New("获取用户关注列表失败")
	}
	return result.value.([]UserFollow), nil
}

// 获取用户粉丝列表
func GetFollowUserList(ctx context.Context, uid int64) ([]UserFollow, error) {
	result := selectList(ctx, int64(0), "user_follow", where{sqlex.Eq{"uid": uid}}, "fuid")
	if !result.success {
		return nil, errors.New("获取用户关注列表失败")
	}
	return result.value.([]UserFollow), nil
}
