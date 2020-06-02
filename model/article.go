package model

import (
	"database/sql"
	"errors"
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
	Updated
)

type Article struct {
	Id        int            `graphql:"id" json:"id"`
	Title     string         `graphql:"title" json:"title"`
	Uid       int            `graphql:"-" json:"uid"`
	Cover     sql.NullString `graphql:"cover;;null" json:"cover,omitempty"`
	SubTitle  string         `graphql:"subTitle" json:"subTitle"`
	Content   string         `graphql:"content" json:"content"`
	State     ArticleState   `graphql:"state" json:"state"`
	CreatedAt time.Time      `graphql:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `graphql:"updatedAt" json:"updatedAt"`
	Author    string         `graphql:"-" json:"author"`
	ArticleEx `graphql:"-" json:"ex"`
}

type ArticleEx struct {
	ViewNum int `json:"viewNum"`
	CmtNum  int `json:"cmtNum"`
	LikeNum int `json:"likeNum"`
}

type ArticleArg struct {
	Id        uint64
	Title     string
	Cover     string
	Tag       []string
	Introduce string
	Content   *string
}

// 新增文章
func InsertArticle(tx *sqlog.DB, setMap map[string]interface{}) (int, error) {
	result, err := PSql.Insert("article").
		SetMap(setMap).
		RunWith(tx).Exec()
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	if id == 0 {
		return 0, errors.New("保存失败")
	}

	result, err = PSql.Insert("article_ex").Columns("aid").Values(id).RunWith(tx).Exec()
	if err != nil {
		return 0, err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return 0, errors.New("保存失败")
	}
	return int(id), nil
}

// 修改文章信息
func UpdateArticle(tx *sqlog.DB, id int, setMap map[string]interface{}) error {
	result, err := PSql.Update("article").
		SetMap(setMap).
		Where(sqlex.Eq{"id": id}).
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

// 查询文章内容
func QueryArticle(tx *sqlog.DB, id int) (Article, error) {
	rows, err := PSql.Select("a.id,a.uid,a.title,a.cover,a.sub_title,a.content,a.state,a.created_at",
		"a.updated_at,e.view_num,e.cmt_num,e.like_num").
		From("article a").
		LeftJoin("article_ex e on a.id=e.aid").
		Where(sqlex.Eq{"a.id": id}).
		Limit(1).
		RunWith(tx).Query()
	if err != nil {
		return Article{}, err
	}
	defer rows.Close()

	var article Article
	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Uid, &article.Title, &article.Cover, &article.SubTitle,
			&article.Content, &article.State, &article.CreatedAt, &article.UpdatedAt, &article.ViewNum,
			&article.CmtNum, &article.LikeNum)
		if err != nil {
			return Article{}, err
		}
	}
	return article, nil
}

// 获取文章列表
func QueryArticles(tx *sqlog.DB, condition string, uid int, draft bool) ([]Article, error) {
	rows, err := PSql.Select("a.id,a.uid,a.title,a.cover,a.sub_title,a.state,a.created_at",
		"a.updated_at,e.view_num,e.cmt_num,e.like_num").
		From("article a").
		LeftJoin("article_ex e on a.id=e.aid").
		WhereExpr(
			sqlex.IF{Condition: condition != "", Sq: sqlex.Or{
				sqlex.Like{"a.title": condition},
				sqlex.Like{"a.content": condition},
			}},
			sqlex.IF{Condition: uid != 0, Sq: sqlex.Eq{"a.uid": uid}},
			sqlex.IF{Condition: !draft, Sq: sqlex.NotEq{"a.state": Draft}},
		).
		OrderBy("a.id desc").
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Uid, &article.Title, &article.Cover, &article.SubTitle,
			&article.State, &article.CreatedAt, &article.UpdatedAt, &article.ViewNum,
			&article.CmtNum, &article.LikeNum)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// 删除文章
func DeleteArticle(tx *sqlog.DB, id int) error {
	result, err := PSql.Delete("article").
		Where(sqlex.Eq{"id": id}).
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
