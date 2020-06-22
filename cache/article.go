package cache

import "strconv"

const (
	Article_Cache    = "Article"
	Article_Ex_Cache = "Article_Ex"
)

type Article struct {
	Id        int
	Condition string
	Uid       int
}

func (a Article) GetCacheKey() string {
	return Article_Cache + "_" + strconv.Itoa(a.Id)
}

func (a Article) GetCachesKey() string {
	if a.Uid != 0 {
		return Article_Cache + "_" + strconv.Itoa(a.Uid)
	} else {
		return Article_Cache + "_" + a.Condition
	}
}

type ArticleEx struct {
	Aid int
}

func (a ArticleEx) GetCacheKey() string {
	return Article_Ex_Cache + "_" + strconv.Itoa(a.Aid)
}

func (a ArticleEx) GetCachesKey() string {
	return Article_Ex_Cache + "_" + strconv.Itoa(a.Aid)
}
