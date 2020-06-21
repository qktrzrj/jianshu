package model

import (
	"errors"
	"github.com/shyptr/plugins/sqlog"
	"github.com/shyptr/sqlex"
)

type Objtype int

const (
	ArticleObj Objtype = iota + 1
	CommentObj
	ReplyObj
)

func LikeList(tx *sqlog.DB, uid int, objtyp Objtype) ([]int, error) {
	rows, err := PSql.Select("objid").
		From("`like`").
		Where(sqlex.Eq{"objtype": objtyp, "uid": uid}).
		RunWith(tx).Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []int
	for rows.Next() {
		var objid int
		if err := rows.Scan(&objid); err != nil {
			return nil, err
		}
		list = append(list, objid)
	}
	return list, nil
}

func HasLike(tx *sqlog.DB, uid, objid int, objtyp Objtype) (bool, error) {
	row, err := PSql.Select("count(*)").
		From("`like`").
		Where(sqlex.Eq{"objtype": objtyp, "objid": objid, "uid": uid}).
		RunWith(tx).Query()
	if err != nil {
		return false, err
	}
	defer row.Close()
	if row.Next() {
		var count int
		err := row.Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return false, nil
}

func Like(tx *sqlog.DB, setMap map[string]interface{}) error {
	result, err := PSql.Insert("`like`").
		SetMap(setMap).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("点赞失败")
	}
	return nil
}

func UnLike(tx *sqlog.DB, objtyp Objtype, objid, uid int) error {
	result, err := PSql.Delete("`like`").
		Where(sqlex.Eq{"objtype": objtyp, "objid": objid, "uid": uid}).
		RunWith(tx).Exec()
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("取消点赞失败")
	}
	return nil
}
