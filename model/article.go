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
	Uid       int            `graphql:"uid" json:"uid"`
	Cover     sql.NullString `graphql:"cover;;null" json:"cover,omitempty"`
	SubTitle  string         `graphql:"subTitle" json:"subTitle"`
	Content   string         `graphql:"content" json:"content"`
	State     ArticleState   `graphql:"state" json:"state"`
	CreatedAt time.Time      `graphql:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `graphql:"updatedAt" json:"updatedAt"`
	User      `graphql:"-" json:"user"`
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
	_, err := PSql.Update("article").
		SetMap(setMap).
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Exec()
	if err != nil {
		return err
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

// 查询文章扩展
func QueryArticleEx(tx *sqlog.DB, id int) (ArticleEx, error) {
	rows, err := PSql.Select("e.view_num,e.cmt_num,e.like_num").
		From("article_ex e").
		Where(sqlex.Eq{"e.aid": id}).
		Limit(1).
		RunWith(tx).Query()
	if err != nil {
		return ArticleEx{}, err
	}
	defer rows.Close()

	var article ArticleEx
	if rows.Next() {
		err := rows.Scan(&article.ViewNum, &article.CmtNum, &article.LikeNum)
		if err != nil {
			return article, err
		}
	}
	return article, nil
}

// 获取文章列表
func QueryArticles(tx *sqlog.DB, uid int, draft bool) ([]Article, error) {
	rows, err := PSql.Select("a.id,a.uid,a.title,a.cover,a.sub_title,a.state,a.updated_at,e.view_num,e.cmt_num,e.like_num").
		From("article a").
		LeftJoin("article_ex e on a.id=e.aid").
		LeftJoin("user u on a.uid = u.id").
		WhereExpr(
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
			&article.State, &article.UpdatedAt, &article.ViewNum, &article.CmtNum, &article.LikeNum)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// 根据idlist获取文章列表
func QueryArticlesByIds(tx *sqlog.DB, ids []int) ([]Article, error) {
	rows, err := PSql.Select("a.id,a.uid,a.title,a.cover,a.sub_title,a.state,a.updated_at,e.view_num,e.cmt_num,e.like_num").
		From("article a").
		LeftJoin("article_ex e on a.id=e.aid").
		LeftJoin("user u on a.uid = u.id").
		Where(sqlex.Eq{"a.id": ids}).
		RunWith(tx).Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Uid, &article.Title, &article.Cover, &article.SubTitle,
			&article.State, &article.UpdatedAt, &article.ViewNum, &article.CmtNum, &article.LikeNum)
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

// 增加/减少 文章浏览数/点赞数/评论数
func AddViewOrLikeOrCmt(tx *sqlog.DB, id, typ int, add bool) error {
	updateBuilder := PSql.Update("article_ex").
		Set("updated_at", time.Now())
	addNum := "+1"
	if !add {
		addNum = "-1"
	}
	if typ == 0 {
		updateBuilder = updateBuilder.DirectSet("view_num=view_num" + addNum)
	} else if typ == 1 {
		updateBuilder = updateBuilder.DirectSet("like_num=like_num" + addNum)
	} else {
		updateBuilder = updateBuilder.DirectSet("cmt_num=cmt_num" + addNum)
	}
	result, err := updateBuilder.Where(sqlex.Eq{"aid": id}).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("修改文章扩展信息失败")
	}
	return nil
}
