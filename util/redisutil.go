package util

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisDb *redis.Client

func RedisInit() *redis.Client {
	if RedisDb != nil {
		status := RedisDb.Ping(context.Background())
		if status.Err() != nil {
			return nil
		}
		return RedisDb
	}
	RedisDb = redis.NewClient(&redis.Options{DB: 0})
	status := RedisDb.Ping(context.Background())
	if status.Err() != nil {
		return nil
	}
	return RedisDb
}
