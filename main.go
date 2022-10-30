package main

import (
	"time"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

func main() {
	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
	// rdb.Del("chat_messages")
	rdb.RPush("cron", time.Now().String())
}
