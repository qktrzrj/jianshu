package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type UserFollow struct {
	Id        int64     `graphql:"id" db:"id"`
	Uid       int64     `graphql:"uid" db:"uid"`
	Fuid      int64     `graphql:"fuid" db:"fuid"`
	CreatedAt time.Time `graphql:"-" db:"created_at"`
	UpdatedAt time.Time `graphql:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `graphql:"deletedAt" db:"deleted_at"`
}

func InsertUserFollow(ctx context.Context, uid, fuid int64) error {
	id, err := idfetcher.NextID()
	if err != nil {
		return err
	}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("user_follow").Columns("id,uid,fuid").Values(id, uid, fuid).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("save follow relation failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("save follow relation failed")
	}
	return nil
}

func RemoveUserFollow(ctx context.Context, uid, fuid int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Update("user_follow").
		Set("deleted_at", time.Now()).
		Where(sqlex.Eq{"uid": uid, "fuid": fuid}).
		RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("remove relation failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("remove relation failed")
	}
	return nil
}

// 获取用户关注列表
func GetUserFollowList(ctx context.Context, fuid int64) ([]int64, error) {
	var follows []int64
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("uid").Where(sqlex.Eq{"fuid": fuid}).RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, errors.New("fetch user follow list failed")
	}
	defer rows.Close()
	for rows.Next() {
		var num int64
		err = rows.Scan(&num)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, errors.New("fetch user follow list failed")
		}
		follows = append(follows, num)
	}
	return follows, nil
}

// 获取用户粉丝列表
func GetFollowUserList(ctx context.Context, uid int64) ([]int64, error) {
	var follows []int64
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("fuid").Where(sqlex.Eq{"uid": uid}).RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, errors.New("fetch user follow list failed")
	}
	defer rows.Close()
	for rows.Next() {
		var num int64
		err = rows.Scan(&num)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, errors.New("fetch user follow list failed")
		}
		follows = append(follows, num)
	}
	return follows, nil
}
