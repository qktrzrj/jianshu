package model

import (
	"database/sql"
	"fmt"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"time"
)

type Gender int

const (
	Man Gender = iota + 1
	Woman
	Unknown
)

type UserState int

const (
	Unsigned UserState = iota + 1
	Normal
	Forbidden
	Freeze
)

type User struct {
	Id        uint64         `graphql:"id"`
	Username  string         `graphql:"username"`
	Email     string         `graphql:"email"`
	Password  string         `graphql:"-"`
	Avatar    string         `graphql:"avatar"`
	Gender    Gender         `graphql:"gender"`
	Introduce sql.NullString `graphql:"introduce"`
	State     UserState      `graphql:"state"`
	Root      bool           `graphql:"root"`
	CreatedAt time.Time      `graphql:"createdAt"`
	UpdatedAt time.Time      `graphql:"updatedAt"`
	DeletedAt sql.NullTime   `graphql:"deletedAt"`
	Count     UserCount      `graphql:"-"`
}

type UserCount struct {
	Uid        uint64       `graphql:"-"`
	FansNum    int          `graphql:"fansNum"`
	FollowNum  int          `graphql:"followNum"`
	ArticleNum int          `graphql:"articleNum"`
	Words      int          `graphql:"words"`
	LikeNum    int          `graphql:"likeNum"`
	CreatedAt  time.Time    `graphql:"-"`
	UpdatedAt  time.Time    `graphql:"-"`
	DeletedAt  sql.NullTime `graphql:"-"`
}

type UserFollow struct {
	Id        uint64       `graphql:"-"`
	Uid       uint64       `graphql:"-"`
	Fuid      uint64       `graphql:"-"`
	CreatedAt time.Time    `graphql:"createdAt"`
	UpdatedAt time.Time    `graphql:"updatedAt"`
	DeletedAt sql.NullTime `graphql:"deletedAt"`
}

func GetUser(tx *sqlog.DB, id uint64, username, email string) (User, error) {
	rows, err := PSql.Select("id,username,email,password,avatar,gender,introduce,state,root,created_at,updated_at,deleted_at").
		From(`"user"`).
		WhereExpr(
			sqlex.IF{id != 0, sqlex.Eq{"id": id}},
		).
		WhereExpr(
			sqlex.Or{
				sqlex.IF{username != "", sqlex.Eq{"username": username}},
				sqlex.IF{email != "", sqlex.Eq{"email": email}},
			},
		).
		RunWith(tx).Query()
	if err != nil {
		return User{}, err
	}
	var user User
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.Gender, &user.Introduce, &user.State,
			&user.Root, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func GetUserCount(tx *sqlog.DB, id uint64) (UserCount, error) {
	rows, err := PSql.Select("fans_num,follow_num,article_num,words,like_num").
		From("user_count").
		Where("uid=$1", id).
		RunWith(tx).Query()
	if err != nil {
		return UserCount{}, err
	}
	var c UserCount
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&c.FansNum, &c.FollowNum, &c.ArticleNum, &c.Words, &c.LikeNum)
		if err != nil {
			return c, err
		}
	}
	return c, nil
}

func GetUserFollower(tx *sqlog.DB, id uint64) ([]uint64, error) {
	rows, err := PSql.Select("fuid").
		From("user_follow").
		Where("uid=$1", id).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}
	var fs []uint64
	defer rows.Close()
	for rows.Next() {
		var f uint64
		err := rows.Scan(&f)
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return fs, nil
}

func GetFollowUser(tx *sqlog.DB, id uint64) ([]uint64, error) {
	rows, err := PSql.Select("uid").
		From("user_follow").
		Where("fuid=$1", id).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}
	var fs []uint64
	defer rows.Close()
	for rows.Next() {
		var f uint64
		err := rows.Scan(&f)
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return fs, nil
}

type UserArg struct {
	Username string `graphql:"username" validate:"min=6,max=16"`
	Email    string `graphql:"email" validate:"email"`
	Password string `graphql:"password" validate:"min=8"`
	Avatar   string `graphql:"-"`
}

func InsertUser(tx *sqlog.DB, arg UserArg) (uint64, error) {
	id, err := IdFetcher.NextID()
	if err != nil {
		return id, err
	}
	result, err := PSql.Insert(`"user"`).
		Columns("id,username,email,password,avatar").
		Values(id, arg.Username, arg.Email, arg.Password, arg.Avatar).
		RunWith(tx).Exec()
	if err != nil {
		return id, err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return id, fmt.Errorf("保存用户信息失败")
	}
	return id, nil
}

func InsertUserCount(tx *sqlog.DB, id uint64) error {
	result, err := PSql.Insert("user_count").Columns("uid").Values(id).RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("保存用户信息失败")
	}
	return nil
}

func InsertUserFollow(tx *sqlog.DB, uid uint64, fuid uint64) error {
	id, err := IdFetcher.NextID()
	if err != nil {
		return err
	}
	result, err := PSql.Insert("user_follow").Columns("id,uid,fuid").Values(id, uid, fuid).RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("关注失败")
	}
	return nil
}

func DeleteUserFollow(tx *sqlog.DB, id uint64, fuid uint64) error {
	result, err := PSql.Update("user_follow").
		Set("deleted_at", time.Now()).
		Where(sqlex.Eq{"uid": id, "fuid": fuid}).
		RunWith(tx).Exec()

	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("取消关注失败")
	}
	return nil
}
