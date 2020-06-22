package handler

import (
	"github.com/shyptr/jianshu/cache"
	"github.com/shyptr/jianshu/model"
)

func registerCache() {
	cache.Relation(cache.User{}, model.User{})
	cache.Relation(cache.Follow{}, int(0))
	cache.Relation(cache.UserCount{}, model.UserCount{})
}
