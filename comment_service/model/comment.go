package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type CommentState int

const (
	Unaudited CommentState = iota + 1
	Online
	Offline
	Deleted
)

var CommentStateString = map[string]CommentState{
	"Unaudited": Unaudited,
	"Online":    Online,
	"Offline":   Offline,
	"Deleted":   Deleted,
}

type Comment struct {
	Id        int64        `graphql:"id" db:"id"`
	Aid       int64        `graphql:"aid" db:"aid"`
	Uid       int64        `graphql:"uid" db:"uid"`
	Content   int64        `graphql:"content" db:"content"`
	ZanNum    int64        `graphql:"zanNum" db:"zan_num"`
	Floor     int64        `graphql:"floor" db:"floor"`
	State     CommentState `graphql:"state" db:"state"`
	CreatedAt time.Time    `graphql:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `graphql:"updatedAt" db:"updated_at"`
	DeletedAt time.Time    `graphql:"deletedAt" db:"deleted_at"`
}

func GetComments(ctx context.Context, aid int64) ([]Comment, error) {
	var comments []Comment
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("id,uid,content,zan_num,floor,state,create_at,deleted_at").
		From("comment").
		Where(sqlex.Eq{"aid": aid}).
		OrderBy("floor desc").
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, errors.New("fetch comments failed")
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Id, &comment.Uid, &comment.Content, &comment.ZanNum, &comment.Floor, &comment.State,
			&comment.CreatedAt, &comment.DeletedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, errors.New("fetch comments failed")
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func InsertComment(ctx context.Context, cv map[string]interface{}) (Comment, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return Comment{}, err
	}
	cv["id"] = int64(id)
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("comment").SetMap(cv).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return Comment{}, errors.New("save comment failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return Comment{}, errors.New("save comment failed")
	}
	var comment Comment
	row := PSql.Select("comment").
		Columns("id,uid,content,zan_num,floor,state,create_at,deleted_at").
		From("comment").
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).QueryRow()
	err = row.Scan(&comment.Id, &comment.Uid, &comment.Content, &comment.ZanNum, &comment.Floor, &comment.State,
		&comment.CreatedAt, &comment.DeletedAt)
	if err != nil {
		logger.Error().Err(err).Send()
		return comment, errors.New("fetch comments failed")
	}
	return comment, nil
}

func RemoveComment(ctx context.Context, id int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Update("comment").
		Set("deleted_at", time.Now()).
		Set("updated_at", time.Now()).
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("remove comment failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("remove comment failed")
	}
	return nil
}
