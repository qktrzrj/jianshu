package model

import (
	"context"
	"errors"
	"github.com/unrotten/sqlex"
	"time"
)

const (
	Unaudited = "Unaudited"
	Online    = "Online"
	Offline   = "offline"
	Deleted   = "deleted"
)

type Article struct {
	Id        int64     `json:"id" db:"id"`
	Sn        string    `json:"sn" db:"sn"`
	Title     string    `json:"title" db:"title"`
	Uid       int64     `json:"uid" db:"uid"`
	Cover     string    `json:"cover" db:"cover"`
	Content   string    `json:"content" db:"content"`
	Tags      []string  `json:"tags" db:"tags"`
	State     string    `json:"state" db:"state"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func GetArticle(ctx context.Context, where where, columns ...string) (Article, error) {
	result := selectOne(ctx, Article{}, "article", where, columns...)
	if !result.success {
		return Article{}, errors.New("获取文章信息失败")
	}
	return result.value.(Article), nil
}

func GetArticles(ctx context.Context, where where, columns ...string) ([]Article, error) {
	result := selectList(ctx, Article{}, "article", where, columns...)
	if !result.success {
		return nil, errors.New("获取文章信息失败")
	}
	return result.value.([]Article), nil
}

func InsertArticle(ctx context.Context, cv cv) (Article, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return Article{}, err
	}
	cv["id"] = int64(id)
	result := insertOne(ctx, Article{}, "article", cv)
	if !result.success {
		return Article{}, errors.New("新增文章失败")
	}
	return result.value.(Article), nil
}

func UpdateArticle(ctx context.Context, cv cv, id int64) error {
	if !update(ctx, "article", cv, where{sqlex.Eq{"id": id}}).success {
		return errors.New("修改文章失败")
	}
	return nil
}

func RemoveArticle(ctx context.Context, id int64) error {
	if !remove(ctx, "article", where{sqlex.Eq{"id": id}}).success {
		return errors.New("删除文章失败")
	}
	return nil
}
