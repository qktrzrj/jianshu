package model

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type ArticleState int

const (
	Draft ArticleState = iota
	Unaudited
	Online
	Offline
	Deleted
)

type Article struct {
	Id        uint64       `graphql:"id" json:"id"`
	SN        string       `graphql:"sn" json:"sn"`
	Title     string       `graphql:"title" json:"title"`
	Uid       uint64       `graphql:"-" json:"uid"`
	Cover     string       `graphql:"cover" json:"cover,omitempty"`
	Tag       []string     `graphql:"tag" json:"tag,omitempty"`
	Introduce string       `graphql:"introduce" json:"introduce"`
	Content   string       `graphql:"content" json:"content"`
	State     ArticleState `graphql:"state" json:"state"`
	CreatedAt time.Time    `graphql:"createdAt" json:"createdAt"`
	UpdatedAt time.Time    `graphql:"updatedAt" json:"updatedAt"`
	DeletedAt sql.NullTime `graphql:"deletedAt" json:"deletedAt,omitempty"`
	Author    string       `graphql:"-" json:"author"`
	ArticleEx `graphql:"-" json:"ex"`
}

type ArticleEx struct {
	Aid       uint64       `json:"-"`
	ViewNum   int          `json:"viewNum"`
	CmtNum    int          `json:"cmtNum"`
	LikeNum   int          `json:"likeNum"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-"`
}

type ArticleDraftArg struct {
	Id        uint64   `graphql:"id;;null"`
	Title     string   `graphql:"title;;null"`
	Cover     string   `graphql:"cover;;null"`
	Tag       []string `graphql:"tag"`
	Content   string   `graphql:"content;;null"`
	Introduce string   `graphql:"-"`
}

type ArticleArg struct {
	Id      uint64   `graphql:"id;;null"`
	Title   string   `graphql:"title"`
	Cover   string   `graphql:"cover;;null"`
	Tag     []string `graphql:"tag"`
	Content string   `graphql:"content"`
}

func InsertArticle(tx *sqlog.DB, arg ArticleDraftArg, uid uint64, state ArticleState) error {
	result, err := PSql.Insert("article").
		Columns("id,sn,title,cover,tag,content,introduce,uid,state").
		Values(arg.Id, arg.Id, arg.Title, arg.Cover, pq.StringArray(arg.Tag), arg.Content, arg.Introduce, uid, state).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("保存失败")
	}

	result, err = PSql.Insert("article_ex").Columns("aid").Values(arg.Id).RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ = result.RowsAffected()
	if affected == 0 {
		return errors.New("保存失败")
	}
	return nil
}

func UpdateArticle(tx *sqlog.DB, arg ArticleDraftArg, state ArticleState) error {
	result, err := PSql.Update("article").
		SetMap(map[string]interface{}{
			"title":      arg.Title,
			"cover":      arg.Cover,
			"tag":        pq.StringArray(arg.Tag),
			"content":    arg.Content,
			"state":      state,
			"introduce":  arg.Introduce,
			"updated_at": time.Now(),
		}).
		Where("id = $1", arg.Id).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("修改失败")
	}
	return nil
}

func QueryArticle(tx *sqlog.DB, id uint64) (Article, error) {
	rows, err := PSql.Select("a.id,a.uid,a.title,a.cover,a.tag,a.introduce,a.content,a.state,a.created_at",
		"a.updated_at,a.deleted_at,e.view_num,e.cmt_num,e.like_num").
		From("article a").
		LeftJoin("article_ex e on a.id=e.aid").
		Where("a.id=$1", id).
		Limit(1).
		RunWith(tx).Query()
	if err != nil {
		return Article{}, err
	}
	defer rows.Close()

	var article Article
	var tag pq.StringArray
	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Uid, &article.Title, &article.Cover, &tag, &article.Introduce,
			&article.Content, &article.State, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt, &article.ViewNum,
			&article.CmtNum, &article.LikeNum)
		if err != nil {
			return Article{}, err
		}
		article.Tag = tag
	}
	return article, nil
}

func QueryArticles(tx *sqlog.DB, condition string, uid uint64) ([]Article, error) {
	rows, err := PSql.Select("a.id,a.uid,a.title,a.cover,a.tag,a.introduce,a.state,a.created_at",
		"a.updated_at,a.deleted_at,e.view_num,e.cmt_num,e.like_num").
		From("article a").
		LeftJoin("article_ex e on a.id=e.aid").
		WhereExpr(
			sqlex.IF{Condition: condition != "", Sq: sqlex.Or{
				sqlex.Like{"title": condition},
				sqlex.Like{"author": condition},
				sqlex.Like{"content": condition},
			}},
			sqlex.IF{Condition: uid != 0, Sq: sqlex.Eq{"uid": uid}},
		).
		Where("deleted_at is null").
		OrderBy("id desc").
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	var tag pq.StringArray
	if rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Uid, &article.Title, &article.Cover, &tag, &article.Introduce,
			&article.State, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt, &article.ViewNum,
			&article.CmtNum, &article.LikeNum)
		if err != nil {
			return nil, err
		}
		article.Tag = tag
		articles = append(articles, article)
	}
	return articles, nil
}

func DeleteArticle(tx *sqlog.DB, id uint64) error {
	result, err := PSql.Update("article").
		Set("deleted_at", time.Now()).
		RunWith(tx).Exec()

	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("删除失败")
	}
	return nil
}
