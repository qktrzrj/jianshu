package cache

import "strconv"

const (
	Comment_Cache = "Comment"
)

type Comment struct {
	Aid int
}

func (c Comment) GetCacheKey() string {
	return Comment_Cache + "_" + strconv.Itoa(c.Aid)
}

func (c Comment) GetCachesKey() string {
	return Comment_Cache + "_" + strconv.Itoa(c.Aid)
}
