package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type CommentReply struct {
	Id        int64     `graphql:"id" db:"id"`
	Cid       int64     `graphql:"cid" db:"cid"`
	Uid       int64     `graphql:"uid" db:"uid"`
	Content   string    `graphql:"content" db:"content"`
	State     string    `graphql:"state" db:"state"`
	CreatedAt time.Time `graphql:"createdAt" db:"created_at"`
	UpdatedAt time.Time `graphql:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `graphql:"deletedAt" db:"deleted_at"`
}

func GetReplies(ctx context.Context, cid int64) ([]CommentReply, error) {
	var replies []CommentReply
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("id,uid,content,state,created_at").
		From("comment_reply").
		Where(sqlex.Eq{"cid": cid}).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, errors.New("fetch replies of comment failed")
	}
	defer rows.Close()
	for rows.Next() {
		var c CommentReply
		err := rows.Scan(&c.Id, &c.Uid, &c.Content, &c.State, &c.CreatedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, errors.New("fetch replies of comment failed")
		}
		replies = append(replies, c)
	}
	return replies, nil
}

func InsertReply(ctx context.Context, cv map[string]interface{}) (CommentReply, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return CommentReply{}, err
	}
	cv["id"] = int64(id)
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("comment_reply").SetMap(cv).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return CommentReply{}, errors.New("save comment reply failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return CommentReply{}, errors.New("save comment reply failed")
	}
	c := CommentReply{}
	rows, err := PSql.Select("id,uid,content,state,created_at").
		From("comment_reply").
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return c, errors.New("fetch replies of comment failed")
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&c.Id, &c.Uid, &c.Content, &c.State, &c.CreatedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return c, errors.New("fetch replies of comment failed")
		}
	}
	return c, nil
}

func RemoveReply(ctx context.Context, id int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Update("comment_reply").
		Set("updated_at", time.Now()).
		Set("deleted_at", time.Now()).
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("delete reply failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("delete reply failed")
	}
	return nil
}
