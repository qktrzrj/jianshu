package model

import (
	"context"
	"errors"
	"github.com/unrotten/builder"
	"github.com/unrotten/sqlex"
	"time"
)

type ArticleEx struct {
	Aid       int64     `json:"aid" db:"aid"`
	ViewNum   int       `json:"viewNum" db:"view_num"`
	CmtNum    int       `json:"cmtNum" db:"cmt_num"`
	ZanNum    int       `json:"zanNum" db:"zan_num"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

func InsertArticleEx(ctx context.Context, cv cv) error {
	if !insertOne(ctx, "article_ex", cv).success {
		return errors.New("保存文章计数信息失败")
	}
	return nil
}

func GetArticleEx(ctx context.Context, aid int64) (ArticleEx, error) {
	result := selectOne(ctx, "article_ex", where{sqlex.Eq{"aid": aid}})
	if !result.success {
		return ArticleEx{}, errors.New("获取文章计数信息失败")
	}
	return builder.GetStructLikeByTag(result.b, ArticleEx{}, "db").(ArticleEx), nil
}

func UpdateArticleEx(ctx context.Context, aid int64, add bool, columns ...string) error {
	directSets, directSet := make([]string, 0, len(columns)), " + 1"
	if !add {
		directSet = " - 1"
	}
	for _, col := range columns {
		directSets = append(directSets, col+directSet)
	}
	if !update(ctx, "article_ex", cv{}, where{sqlex.Eq{"aid": aid}}, directSets...).success {
		return errors.New("修改文章计数失败")
	}
	return nil
}
