package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/shyptr/jianshu/setting"
	"log"
	"sync"
	"time"
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

// Set a key/value
func Set(key string, data interface{}, time time.Duration) error {
	conn := RedisClient.Conn()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Set(key, value, time).Result()
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisClient.Conn()
	defer conn.Close()

	exists, err := conn.Exists(key).Result()
	if err != nil {
		return false
	}

	return exists > 0
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisClient.Conn()
	defer conn.Close()

	reply, err := conn.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(reply), nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisClient.Conn()
	defer conn.Close()

	result, err := conn.Del(key).Result()
	return result > 0, err
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisClient.Conn()
	defer conn.Close()

	keys, err := conn.Keys("*" + key + "*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
