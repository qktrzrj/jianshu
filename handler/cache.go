package handler

import (
	"github.com/shyptr/jianshu/cache"
	"github.com/shyptr/jianshu/model"
)

func registerCache() {
	cache.Relation(cache.User{}, model.User{})
	cache.Relation(cache.UserCount{}, model.UserCount{})
	cache.Relation(cache.Follow{}, int(0))
	cache.Relation(cache.ArticleEx{}, model.ArticleEx{})
	cache.Relation(cache.Article{}, model.Article{})
	cache.Relation(cache.Like{}, int(0))
	cache.Relation(cache.Comment{}, model.Comment{})
	cache.Relation(cache.MsgNum{}, model.MsgNum{})
	cache.Relation(cache.Msg{}, model.Msg{})
	cache.Relation(cache.Reply{}, model.Reply{})
}
