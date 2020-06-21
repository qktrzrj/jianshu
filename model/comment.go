package model

import (
	"errors"
	"fmt"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type Comment struct {
	Id        int       `graphql:"id"`
	Aid       int       `graphql:"-"`
	Uid       int       `graphql:"-"`
	Content   string    `graphql:"content"`
	LikeNum   int       `graphql:"likeNum"`
	Floor     int       `graphql:"floor"`
	UpdatedAt time.Time `graphql:"updatedAt"`
}

func AddComment(tx *sqlog.DB, setMap map[string]interface{}) (Comment, error) {
	result, err := PSql.Insert("`comment`").
		Columns("aid,uid,content,floor").
		Select(
			PSql.Select(fmt.Sprintf("%d as aid,%d as uid,\"%s\" as content,CASE WHEN MAX(floor)>0 THEN MAX(floor)+1 ELSE 1 END as floor",
				setMap["aid"], setMap["uid"], setMap["content"])).
				From("`comment`").
				Where(sqlex.Eq{"aid": setMap["aid"]}),
		).
		RunWith(tx).Exec()
	if err != nil {
		return Comment{}, err
	}
	id, _ := result.LastInsertId()
	if id == 0 {
		return Comment{}, errors.New("添加评论失败")
	}

	comment := Comment{
		Id:      int(id),
		Aid:     setMap["aid"].(int),
		Uid:     setMap["uid"].(int),
		Content: setMap["content"].(string),
	}
	queryRow := PSql.Select("floor,updated_at").
		From("`comment`").
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).QueryRow()
	err = queryRow.Scan(&comment.Floor, &comment.UpdatedAt)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func CommentList(tx *sqlog.DB, aid int) ([]Comment, error) {
	rows, err := PSql.Select("id,aid,uid,content,like_num,floor,updated_at").
		From("comment").
		Where(sqlex.Eq{"aid": aid}).
		OrderBy("id desc").
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}

	var comments []Comment
	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.Id, &c.Aid, &c.Uid, &c.Content, &c.LikeNum, &c.Floor, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
