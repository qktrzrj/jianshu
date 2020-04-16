package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/shyptr/hello-world-web/setting"
)

var Redis *redis.Client

func InitRedis() {
	redisConfig := setting.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisConfig.GetHost(), redisConfig.GetPort()),
		DB:   0,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
