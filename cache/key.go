package cache

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
	"reflect"
	"time"
)

type key interface {
	GetCacheKey() string
	GetCachesKey() string
}

var relation = map[reflect.Type]reflect.Type{}

func Relation(key interface{}, model interface{}) {
	relation[reflect.TypeOf(key)] = reflect.TypeOf(model)
}

func QueryCaches(ctx context.Context, key key, noExist func() (interface{}, error)) (interface{}, error) {
	r := reflect.TypeOf(key)
	logger := ctx.Value("logger").(zerolog.Logger)

	s := reflect.MakeSlice(reflect.SliceOf(relation[r]), 0, 0).Interface()

	// 查询缓存
	k := key.GetCacheKey()
	if Exists(k) {
		result, err := Get(k)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
		} else {
			json.Unmarshal(result, &s)
			return s, nil
		}
	}
	data, err := noExist()
	if err != nil {
		return nil, err
	}
	Set(k, data, time.Hour*2)
	return data, nil
}

func QueryCache(ctx context.Context, key key, noExist func() (interface{}, error)) (interface{}, error) {
	r := reflect.TypeOf(key)
	logger := ctx.Value("logger").(zerolog.Logger)

	s := reflect.New(relation[r]).Elem().Interface()

	// 查询缓存
	k := key.GetCachesKey()
	if Exists(k) {
		result, err := Get(k)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
		} else {
			json.Unmarshal(result, &s)
			return s, nil
		}
	}
	data, err := noExist()
	if err != nil {
		return nil, err
	}
	Set(k, data, time.Hour*2)
	return data, nil
}
