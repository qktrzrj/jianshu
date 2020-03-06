package model

import (
	"context"
	"errors"
	"github.com/unrotten/builder"
	"github.com/unrotten/sqlex"
	"time"
)

type CommentReply struct {
	Id        int64     `json:"id" db:"id"`
	Cid       int64     `json:"cid" db:"cid"`
	Uid       int64     `json:"uid" db:"uid"`
	Content   string    `json:"content" db:"content"`
	State     string    `json:"state" db:"state"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetReplies(ctx context.Context, cid int64) ([]CommentReply, error) {
	result := selectList(ctx, "comment_reply", where{sqlex.Eq{"cid": cid}})
	if !result.success {
		return nil, errors.New("获取评论回复失败")
	}
	list, _ := builder.Get(result.b, "list")
	replies := make([]CommentReply, 0, len(list.([]interface{})))
	for _, item := range list.([]interface{}) {
		replies = append(replies, builder.GetStructLikeByTag(item.(builder.Builder), CommentReply{}, "db").(CommentReply))
	}
	return replies, nil
}

func InsertReply(ctx context.Context, cv cv) (CommentReply, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return CommentReply{}, err
	}
	cv["id"] = int64(id)
	result := insertOne(ctx, "comment_reply", cv)
	if !result.success {
		return CommentReply{}, errors.New("保存回复失败")
	}
	return builder.GetStructLikeByTag(result.b, CommentReply{}, "db").(CommentReply), nil
}

func RemoveReply(ctx context.Context, id int64) error {
	if !remove(ctx, "comment_reply", where{sqlex.Eq{"id": id}}).success {
		return errors.New("删除回复失败")
	}
	return nil
}
