package cache

import (
	"github.com/shyptr/jianshu/model"
	"strconv"
)

const (
	Like_Cache = "Like"
)

type Like struct {
	Uid int
	Typ model.Objtype
}

func (l Like) GetCacheKey() string {
	return Like_Cache + "_" + strconv.Itoa(l.Uid) + "_" + strconv.Itoa(int(l.Typ))
}

func (l Like) GetCachesKey() string {
	return Like_Cache + "_" + strconv.Itoa(l.Uid) + "_" + strconv.Itoa(int(l.Typ))
}
