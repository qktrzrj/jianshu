package model

import (
	"context"
	"errors"
	"github.com/unrotten/builder"
	"github.com/unrotten/sqlex"
	"time"
)

type Comment struct {
	Id        int64     `json:"id" db:"id"`
	Aid       int64     `json:"aid" db:"aid"`
	Uid       int64     `json:"uid" db:"uid"`
	Content   int64     `json:"content" db:"content"`
	ZanNum    int64     `json:"zanNum" db:"zan_num"`
	Floor     int64     `json:"floor" db:"floor"`
	State     string    `json:"state" db:"state"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetComments(ctx context.Context, aid int64) ([]Comment, error) {
	result := selectList(ctx, "comment", where{sqlex.Eq{"aid": aid}})
	if !result.success {
		return nil, errors.New("获取评论失败")
	}
	list, _ := builder.Get(result.b, "list")
	comments := make([]Comment, 0, len(list.([]interface{})))
	for _, item := range list.([]interface{}) {
		comments = append(comments, builder.GetStructLikeByTag(item.(builder.Builder), Comment{}, "db").(Comment))
	}
	return comments, nil
}

func InsertComment(ctx context.Context, cv cv) (Comment, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return Comment{}, err
	}
	cv["id"] = int64(id)
	result := insertOne(ctx, "comment", cv)
	if !result.success {
		return Comment{}, errors.New("保存评论失败")
	}
	return builder.GetStructLikeByTag(result.b, Comment{}, "db").(Comment), nil
}

func UpdateComment(ctx context.Context, id int64, cv cv, add bool, columns ...string) error {
	var directSets []string
	if len(columns) > 0 {
		if add {
			directSets = append(directSets, "zan_num + 1")
		} else {
			directSets = append(directSets, "zan_num - 1")
		}
	}
	if !update(ctx, "comment", cv, where{sqlex.Eq{"id": id}}, directSets...).success {
		return errors.New("修改评论失败")
	}
	return nil
}

func RemoveComment(ctx context.Context, id int64) error {
	if !remove(ctx, "comment", where{sqlex.Eq{"id": id}}).success {
		return errors.New("删除评论失败")
	}
	return nil
}
