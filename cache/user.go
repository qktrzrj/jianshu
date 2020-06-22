package cache

import (
	"strconv"
	"strings"
)

const (
	User_Cache       = "User"
	Follow_Cache     = "Follow"
	User_Count_Cache = "User_Count"
)

type User struct {
	Id       int
	Username string
}

func (u User) GetCacheKey() string {
	return User_Cache + "_" + strconv.Itoa(u.Id)
}

func (u User) GetCachesKey() string {
	keys := []string{
		User_Cache,
		"LIST",
	}
	if u.Id > 0 {
		keys = append(keys, strconv.Itoa(u.Id))
	}
	if u.Username != "" {
		keys = append(keys, u.Username)
	}
	return strings.Join(keys, "_")
}

type Follow struct {
	Uid  int
	Fuid int
}

func (f Follow) GetCacheKey() string {
	return Follow_Cache + "_" + strconv.Itoa(f.Uid) + "_" + strconv.Itoa(f.Fuid)
}

func (f Follow) GetCachesKey() string {
	keys := []string{
		Follow_Cache,
		"LIST",
	}
	if f.Uid > 0 {
		keys = append(keys, strconv.Itoa(f.Uid))
	} else if f.Fuid > 0 {
		keys = append(keys, strconv.Itoa(f.Fuid))
	}
	return strings.Join(keys, "_")
}

type UserCount struct {
	Uid int
}

func (u UserCount) GetCacheKey() string {
	return User_Count_Cache + "_" + strconv.Itoa(u.Uid)
}

func (u UserCount) GetCachesKey() string {
	return User_Count_Cache + "_" + strconv.Itoa(u.Uid)
}
