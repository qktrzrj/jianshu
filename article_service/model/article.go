package model

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type ArticleState int

const (
	Unaudited ArticleState = iota + 1
	Online
	Offline
	Deleted
)

var ArticleStateString = map[string]ArticleState{
	"Unaudited": Unaudited,
	"Online":    Online,
	"Offline":   Offline,
	"Deleted":   Deleted,
}

type Article struct {
	Id        int64        `graphql:"id" db:"id"`
	Sn        string       `graphql:"sn" db:"sn"`
	Title     string       `graphql:"title" db:"title"`
	Uid       int64        `graphql:"uid" db:"uid"`
	Cover     *string      `graphql:"cover" db:"cover"`
	Content   string       `graphql:"content" db:"content"`
	Tags      []string     `graphql:"tags" db:"tags"`
	State     ArticleState `graphql:"state" db:"state"`
	CreatedAt time.Time    `graphql:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `graphql:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time   `graphql:"deletedAt" db:"deleted_at"`
}

func GetArticle(ctx context.Context, where ...sqlex.Sqlex) (Article, error) {
	article := Article{}
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("id,sn,title,uid,cover,content,tags,state,created_at,updated_at,deleted_at").
		From("article").
		WhereExpr(where...).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return Article{}, errors.New("fetch article content failed")
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&article.Id, &article.Sn, &article.Title, &article.Uid, &article.Cover, &article.Tags, &article.State,
			&article.CreatedAt, &article.UpdatedAt, &article.DeletedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return Article{}, errors.New("fetch article content failed")
		}
	}
	return article, nil
}

func GetArticles(ctx context.Context, where ...sqlex.Sqlex) ([]Article, error) {
	var articles []Article
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	rows, err := PSql.Select("id,sn,title,uid,cover,content,tags,state,created_at,updated_at,deleted_at").
		From("article").
		WhereExpr(where...).
		RunWith(tx).Query()
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, errors.New("fetch article list failed")
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Id, &article.Sn, &article.Title, &article.Uid, &article.Cover, &article.Tags, &article.State,
			&article.CreatedAt, &article.UpdatedAt, &article.DeletedAt)
		if err != nil {
			logger.Error().Err(err).Send()
			return nil, errors.New("fetch article list failed")
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func InsertArticle(ctx context.Context, cv map[string]interface{}) (Article, error) {
	id, err := idfetcher.NextID()
	if err != nil {
		return Article{}, err
	}
	cv["id"] = int64(id)
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Insert("article").SetMap(cv).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return Article{}, errors.New("create new article failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return Article{}, errors.New("create new article failed")
	}
	return GetArticle(ctx, sqlex.Eq{"id": id})
}

func UpdateArticle(ctx context.Context, cv map[string]interface{}, id int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	cv["updated_at"] = time.Now()
	result, err := PSql.Update("article").SetMap(cv).Where(sqlex.Eq{"id": id}).RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("update article content failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("update article content failed")
	}
	return nil
}

func RemoveArticle(ctx context.Context, id int64) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)
	result, err := PSql.Update("article").
		Set("updated_at", time.Now()).
		Set("deleted_at", time.Now()).
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Exec()
	if err != nil {
		logger.Error().Err(err).Send()
		return errors.New("delete article failed")
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("delete article failed")
	}
	return nil
}
