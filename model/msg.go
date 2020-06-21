package model

import (
	"fmt"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type MsgType int

const (
	CommentMsg MsgType = iota + 1
	ReplyMsg
	LikeMsg
	FollowMsg
)

type MsgNum struct {
	Comment int `graphql:"comment"`
	Like    int `graphql:"like"`
	Follow  int `graphql:"follow"`
	Reply   int `graphql:"reply"`
}

type Msg struct {
	Id        int       `graphql:"id"`
	FromId    int       `graphql:"-"`
	Type      MsgType   `graphql:"type"`
	Content   string    `graphql:"content"`
	UpdatedAt time.Time `graphql:"updatedAt"`
}

func AddMsg(tx *sqlog.DB, typ MsgType, from, to int, content string) error {
	_, err := PSql.Insert("`msg`").Columns("typ,from_id,to_id,content").Values(typ, from, to, content).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	return nil
}

func ReadMsg(tx *sqlog.DB, uid int, typ MsgType) error {
	_, err := PSql.Update("`msg`").Set("`read`", true).Where(sqlex.Eq{"typ": typ, "to_id": uid}).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	return nil
}

func ListMsg(tx *sqlog.DB, typ MsgType, uid int) ([]Msg, error) {
	rows, err := PSql.Select("id,from_id,typ,content,updated_at").
		From("`msg`").
		Where(sqlex.Eq{"to_id": uid, "typ": typ}).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var msgs []Msg
	for rows.Next() {
		var msg Msg
		err := rows.Scan(&msg.Id, &msg.FromId, &msg.Type, &msg.Content, &msg.UpdatedAt)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

func QueryMsgNum(tx *sqlog.DB, uid int) (MsgNum, error) {
	rows, err := PSql.Select(fmt.Sprintf("(SELECT COUNT(*) FROM msg WHERE `read`=0 and typ=1 and to_id=%d) as cn,"+
		"(SELECT COUNT(*) FROM msg WHERE `read`=0 and typ=2 and to_id=%d) as rn,"+
		"(SELECT COUNT(*) FROM msg WHERE `read`=0 and typ=3 and to_id=%d) as ln,"+
		"(SELECT COUNT(*) FROM msg WHERE `read`=0 and typ=4 and to_id=%d) as fn", uid, uid, uid, uid)).
		RunWith(tx).Query()
	if err != nil {
		return MsgNum{}, err
	}
	defer rows.Close()

	var msgNum MsgNum
	if rows.Next() {
		err := rows.Scan(&msgNum.Comment, &msgNum.Reply, &msgNum.Like, &msgNum.Follow)
		if err != nil {
			return msgNum, err
		}
	}
	return msgNum, nil
}
