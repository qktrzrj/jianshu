package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type ObjType int

const (
	ArticleObj ObjType = iota + 1
	CommentObj
	ReplyObj
)

var ObjTypeString = map[string]ObjType{
	"Article": ArticleObj,
	"Comment": CommentObj,
	"Reply":   ReplyObj,
}

type Zan struct {
	Id        int64     `graphql:"id" db:"id"`
	Uid       int64     `graphql:"uid" db:"uid"`
	Objtype   ObjType   `graphql:"objtype" db:"objtype"`
	Objid     int64     `graphql:"objid" db:"objid"`
	CreatedAt time.Time `graphql:"-" db:"created_at"`
	UpdatedAt time.Time `graphql:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `graphql:"deletedAt" db:"deleted_at"`
}

func DeleteZan(ctx context.Context, id int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Update("zan").
		Set("updated_at", time.Now()).
		Set("deleted_at", time.Now()).
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("remove zan data failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("remove zan data failed")
	}
	return nil
}

func InsertZan(ctx context.Context, cv map[string]interface{}) (Zan, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return Zan{}, err
	}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	cv["id"] = int64(id)
	result, err := PSql.Insert("zan").SetMap(cv).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return Zan{}, errors.New("save zan data failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return Zan{}, errors.New("save zan data failed")
	}
	var zan Zan
	rows, err := PSql.Select("id,uid,objtype,objid,created_at,deleted_at").
		From("zan").
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return Zan{}, errors.New("save zan data failed")
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&zan.Id, &zan.Uid, &zan.Objtype, &zan.Objid, &zan.CreatedAt, &zan.DeletedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return Zan{}, errors.New("save zan data failed")
		}
	}
	return zan, nil
}
