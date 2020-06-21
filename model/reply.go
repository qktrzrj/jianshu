package model

import (
	"errors"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type Reply struct {
	Id        int       `graphql:"id"`
	Cid       int       `graphql:"-"`
	Uid       int       `graphql:"-"`
	Content   string    `graphql:"content"`
	UpdatedAt time.Time `graphql:"updatedAt"`
}

func ListReply(tx *sqlog.DB, cid int) ([]Reply, error) {
	rows, err := PSql.Select("id,cid,uid,content,updated_at").
		From("comment_reply").
		Where(sqlex.Eq{"cid": cid}).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var replies []Reply
	for rows.Next() {
		var reply Reply
		err := rows.Scan(&reply.Id, &reply.Cid, &reply.Uid, &reply.Content, &reply.UpdatedAt)
		if err != nil {
			return nil, err
		}
		replies = append(replies, reply)
	}
	return replies, nil
}

func AddReply(tx *sqlog.DB, setMap map[string]interface{}) (Reply, error) {
	result, err := PSql.Insert("comment_reply").
		SetMap(setMap).
		RunWith(tx).Exec()
	if err != nil {
		return Reply{}, err
	}

	id, _ := result.LastInsertId()
	if id == 0 {
		return Reply{}, errors.New("添加回复失败")
	}

	row := PSql.Select("id,cid,uid,content,updated_at").
		From("comment_reply").
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).QueryRow()
	var reply Reply
	err = row.Scan(&reply.Id, &reply.Cid, &reply.Uid, &reply.Content, &reply.UpdatedAt)
	if err != nil {
		return Reply{}, err
	}
	return reply, nil
}
