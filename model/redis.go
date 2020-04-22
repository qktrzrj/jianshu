package model

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/shyptr/jianshu/setting"
	"log"
	"sync"
)

var (
	RedisClient *redis.Client
	rOnce       sync.Once
)

func RedisInit() {
	rOnce.Do(func() {
		config := setting.GetCache()
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.GetHost(), config.GetPort()),
			Password: config.GetPassword(),
			DB:       config.GetDB(),
		})

		_, err := RedisClient.Ping().Result()
		if err != nil {
			log.Fatalf("连接redis失败: %s", err)
		}
	})
}
