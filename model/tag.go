package model

import (
	"context"
	"errors"
	"github.com/unrotten/builder"
	"github.com/unrotten/sqlex"
	"time"
)

type Tag struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func InsertTag(ctx context.Context, name string) error {
	id, err := idfetcher.NextID()
	if err != nil {
		return err
	}
	if !insertOne(ctx, "tag", cv{"id": int64(id), "name": name}).success {
		return errors.New("新增标签失败")
	}
	return nil
}

func GetTags(ctx context.Context, name string) ([]Tag, error) {
	result := selectList(ctx, "tag", where{sqlex.IF{Condition: name != "", Sq: sqlex.Like{"name": name}}})
	if !result.success {
		return nil, errors.New("获取标签列表失败")
	}
	list, _ := builder.Get(result.b, "list")
	tags := make([]Tag, 0, len(list.([]interface{})))
	for _, item := range list.([]interface{}) {
		tags = append(tags, builder.GetStructLikeByTag(item.(builder.Builder), Tag{}, "db").(Tag))
	}
	return tags, nil
}
