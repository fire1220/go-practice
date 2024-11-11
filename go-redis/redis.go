package go_redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisLock struct {
}

func (r *RedisLock) Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.17.223.134:6379",
		Password: "",
		DB:       3,
	})
	return rdb
}

func (r *RedisLock) Lock(ctx context.Context, key, val string, expire time.Duration) bool {
	if key == "" || val == "" {
		return false
	}
	return r.Connect().SetNX(ctx, key, val, expire).Val()
}

func (r *RedisLock) Unlock(ctx context.Context, key, val string) error {
	script := `
            if redis.call("get", KEYS[1]) == ARGV[1] then
                return redis.call("del", KEYS[1])
            else
                return 0
            end
`
	res, err := r.Connect().Eval(ctx, script, []string{key}, val).Int()
	if err != nil {
		return err
	}
	if res == 0 {
		return errors.New("unlocking failed")
	}
	return nil
}
