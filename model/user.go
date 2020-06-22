package model

import (
	"database/sql"
	"fmt"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
	"strconv"
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
	Id        int            `graphql:"id"`
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
	Count     UserCount      `graphql:"-"`
}

type UserCount struct {
	FansNum    int `graphql:"fansNum"`
	FollowNum  int `graphql:"followNum"`
	ArticleNum int `graphql:"articleNum"`
	Words      int `graphql:"words"`
	LikeNum    int `graphql:"likeNum"`
}

// 获取用户基本信息
// 根据id，username和email查询，若无过滤条件，默认取第一条
func GetUser(tx *sqlog.DB, id int, username, email string) (User, error) {
	rows, err := PSql.Select("id,username,email,password,avatar,gender,introduce,state,root,created_at,updated_at").
		From("`user`").
		Where("1=1").
		WhereExpr(
			sqlex.IF{Condition: id != 0, Sq: sqlex.Eq{"id": id}},
			sqlex.Or{
				sqlex.IF{Condition: username != "", Sq: sqlex.Eq{"username": username}},
				sqlex.IF{Condition: email != "", Sq: sqlex.Eq{"email": email}},
			},
		).
		Limit(1).
		RunWith(tx).Query()
	if err != nil {
		return User{}, err
	}
	var user User
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.Gender,
			&user.Introduce, &user.State, &user.Root, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

// 用户列表
func GetUsers(tx *sqlog.DB, username string) ([]User, error) {
	rows, err := PSql.Select("u.id,u.username,u.avatar,c.fans_num,c.follow_num,c.article_num,c.words,c.like_num").
		From("`user` u").
		LeftJoin("user_count c on c.uid=u.id").
		Where("1=1").
		Where(sqlex.IF{Condition: username != "", Sq: sqlex.Like{"username": username + "%"}}).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}

	var users []User

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Avatar, &user.Count.FansNum, &user.Count.FollowNum,
			&user.Count.ArticleNum, &user.Count.Words, &user.Count.LikeNum)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// 获取用户扩展信息
func GetUserCount(tx *sqlog.DB, id int) (UserCount, error) {
	rows, err := PSql.Select("fans_num,follow_num,article_num,words,like_num").
		From("user_count").
		Where(sqlex.Eq{"uid": id}).
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

// 获取用户粉丝列表
func GetUserFollower(tx *sqlog.DB, id int) ([]int, error) {
	rows, err := PSql.Select("fuid").
		From("user_follow").
		Where(sqlex.Eq{"uid": id}).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}
	var fs []int
	defer rows.Close()
	for rows.Next() {
		var f int
		err := rows.Scan(&f)
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return fs, nil
}

// 获取用户关注列表
func GetFollowUser(tx *sqlog.DB, id int) ([]int, error) {
	rows, err := PSql.Select("uid").
		From("user_follow").
		Where(sqlex.Eq{"fuid": id}).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}
	var fs []int
	defer rows.Close()
	for rows.Next() {
		var f int
		err := rows.Scan(&f)
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return fs, nil
}

// 新增用户
func InsertUser(tx *sqlog.DB, arg map[string]interface{}) (int, error) {
	// 新增用户基本信息
	result, err := PSql.Insert("`user`").
		SetMap(arg).
		RunWith(tx).Exec()
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	if id == 0 {
		return 0, fmt.Errorf("保存用户信息失败")
	}
	// 用户拓展信息
	result, err = PSql.Insert("user_count").Columns("uid").Values(id).RunWith(tx).Exec()
	if err != nil {
		return 0, err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return 0, fmt.Errorf("保存用户信息失败")
	}
	return int(id), nil
}

// 关注
func InsertUserFollow(tx *sqlog.DB, uid, fuid int) error {
	result, err := PSql.Insert("user_follow").Columns("uid,fuid").Values(uid, fuid).RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("关注失败")
	}
	return nil
}

// 取消关注
func DeleteUserFollow(tx *sqlog.DB, id, fuid int) error {
	result, err := PSql.Delete("user_follow").
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

// 修改用户基本信息
func UpdateUser(tx *sqlog.DB, id int, setMap map[string]interface{}) error {
	_, err := PSql.Update("`user`").
		SetMap(setMap).
		Where(sqlex.Eq{"id": id}).
		RunWith(tx).Exec()

	if err != nil {
		return err
	}
	return nil
}

// 用户关系
func IsFollow(tx *sqlog.DB, uid, fuid int) (bool, error) {
	row := PSql.Select("count(*)").
		From("user_follow").
		Where(sqlex.Eq{"uid": uid, "fuid": fuid}).
		RunWith(tx).QueryRow()
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// 修改计数
func UpdateUserCount(tx *sqlog.DB, uid, typ int, add bool, word ...int) error {
	count := "+1"
	if !add {
		count = "-1"
	}
	updateBuilder := PSql.Update("user_count").Set("updated_at", time.Now()).Where(sqlex.Eq{"uid": uid})
	// 粉丝数
	if typ == 0 {
		updateBuilder.DirectSet("fans_num=fans_num" + count)
	}
	// 关注数
	if typ == 1 {
		updateBuilder.DirectSet("follow_num=fans_num" + count)
	}
	// 文章数
	if typ == 2 {
		updateBuilder.DirectSet("article_num=article_num" + count)
	}
	// 字数
	if typ == 3 {
		count = "+"
		if !add {
			count = "-"
		}
		updateBuilder.DirectSet("words=words" + count + strconv.Itoa(word[0]))
	}
	// 赞
	if typ == 4 {
		updateBuilder.DirectSet("like_num=like_num" + count)
	}
	_, err := updateBuilder.RunWith(tx).Exec()
	return err
}
