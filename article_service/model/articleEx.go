package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"time"
)

type ArticleEx struct {
	Aid       int64     `graphql:"-" db:"aid"`
	ViewNum   int       `graphql:"viewNum" db:"view_num"`
	CmtNum    int       `graphql:"cmtNum" db:"cmt_num"`
	ZanNum    int       `graphql:"zanNum" db:"zan_num"`
	CreatedAt time.Time `graphql:"-" db:"created_at"`
	UpdatedAt time.Time `graphql:"-" db:"updated_at"`
	DeletedAt time.Time `graphql:"-" db:"deleted_at"`
}

func InsertArticleEx(ctx context.Context, aid int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("article_ex").Columns("aid").Values(aid).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("save article info failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("save article info failed")
	}
	return nil
}

func GetArticleEx(ctx context.Context, aid int64) (ArticleEx, error) {
	articleEx := ArticleEx{}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("article_ex").
		Columns("view_num,cmt_num,zan_num").
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return articleEx, errors.New("get article info failed")
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&articleEx.ViewNum, &articleEx.CmtNum, &articleEx.ZanNum)
		if err != nil {
			return articleEx, errors.New("get article info failed")
		}
	}
	return articleEx, nil
}
