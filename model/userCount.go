package model

import (
	"context"
	"errors"
	"github.com/unrotten/builder"
	"github.com/unrotten/sqlex"
	"time"
)

type UserCount struct {
	Uid        int64     `json:"uid" db:"uid"`
	FansNum    int32     `json:"fansNum" db:"fans_num"`
	FollowNum  int32     `json:"followNum" db:"follow_num"`
	ArticleNum int32     `json:"articleNum" db:"article_num"`
	Words      int32     `json:"words" db:"words"`
	ZanNum     int32     `json:"zanNum" db:"zan_num"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt  time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetUserCount(ctx context.Context, uid int64) (UserCount, error) {
	result := selectOne(ctx, "user_count", append(where{}, sqlex.Eq{"uid": uid}))
	if !result.success {
		return UserCount{}, errors.New("查询用户计数失败")
	}
	return builder.GetStructLikeByTag(result.b, UserCount{}, "db").(UserCount), nil
}

func InsertUserCount(ctx context.Context, uid int64) error {
	result := insertOne(ctx, "user_count", cv{"uid": uid})
	if !result.success {
		return errors.New("保存用户计数表失败")
	}
	return nil
}

func UpdateUserCount(ctx context.Context, uid int64, add bool, columns ...string) error {
	directSets, directSet := make([]string, 0, len(columns)), " + 1"
	if !add {
		directSet = " - 1"
	}
	for _, col := range columns {
		directSets = append(directSets, col+directSet)
	}
	if !update(ctx, "user_count", cv{}, where{sqlex.Eq{"uid": uid}}, directSets...).success {
		return errors.New("修改用户计数失败")
	}
	return nil
}
