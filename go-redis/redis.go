package go_redis

import "github.com/go-redis/redis/v8"

func Connection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "172.17.223.134:6379",
		Password: "",
		DB:       3,
	})
}
