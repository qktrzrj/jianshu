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

type UserCount struct {
	Uid        int64      `graphql:"-" db:"uid"`
	FansNum    int32      `graphql:"fansNum" db:"fans_num"`
	FollowNum  int32      `graphql:"followNum" db:"follow_num"`
	ArticleNum int32      `graphql:"articleNum" db:"article_num"`
	Words      int32      `graphql:"words" db:"words"`
	ZanNum     int32      `graphql:"zanNum" db:"zan_num"`
	CreatedAt  time.Time  `graphql:"-" db:"created_at"`
	UpdatedAt  time.Time  `graphql:"updatedAt" db:"updated_at"`
	DeletedAt  *time.Time `graphql:"deletedAt" db:"deleted_at"`
}

func GetUserCount(ctx context.Context, uid int64) (UserCount, error) {
	userCount := UserCount{}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("fans_num,follow_num,article_num,words,zan_num").
		From("user_count").
		Where(sqlex.Eq{"uid": uid}).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return UserCount{}, fmt.Errorf("fetch user count info failed")
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&userCount.FansNum, &userCount.FollowNum, &userCount.ArticleNum, &userCount.Words, &userCount.ZanNum)
		if err != nil {
			logger.Error().Err(err).Send()
			return UserCount{}, fmt.Errorf("fetch user count info failed")
		}
	}
	return userCount, nil
}

func InsertUserCount(ctx context.Context, uid int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("user_count").Columns("uid").Values(uid).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return fmt.Errorf("save user count info failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("save user count info failed")
	}
	return nil
}
