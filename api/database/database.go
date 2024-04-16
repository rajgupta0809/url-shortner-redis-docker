package database

import (
	"context"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       dbNo,
	})

	return rdb
}
