package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

func main() {

	redisSentAddrs := "sentialXXX.net:port,sentinalXXX.net:port,sentinalXXX.net:port"
	redisOpts := &redis.FailoverOptions{
		MasterName:       "mymaster",
		SentinelAddrs:    strings.Split(redisSentAddrs, ","),
		SentinelPassword: "sent-pass",
		Password:         "pass",
		DB:               0,
	}

	rdb := redis.NewFailoverClient(redisOpts)

	fmt.Println(rdb.Ping(context.Background()))

	key, err := rdb.RPop(context.Background(), "queue_name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key)
}
