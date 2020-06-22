package cache

import "strconv"

const (
	Reply_Cache = "Reply"
)

type Reply struct {
	Cid int
}

func (r Reply) GetCacheKey() string {
	return Reply_Cache + "_" + strconv.Itoa(r.Cid)
}

func (r Reply) GetCachesKey() string {
	return Reply_Cache + "_" + strconv.Itoa(r.Cid)
}
