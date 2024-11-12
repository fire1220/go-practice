package go_redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

type lock struct {
	redis *redis.Client
	key   string
	val   string
}

var (
	ErrNoLock = errors.New("no lock")
	ErrNoData = errors.New("no data to unlock")
)

func NewLock() *lock {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.17.223.134:6379",
		Password: "",
		DB:       3,
	})
	n := new(lock)
	n.redis = rdb
	return n
}

func (l *lock) Lock(ctx context.Context, key string, expire time.Duration) error {
	if key == "" {
		return errors.New("redis key cannot be empty")
	}
	val, err := l.LockString(ctx, key, expire)
	l.key = key
	l.val = val
	return err
}

func (l *lock) Unlock(ctx context.Context) error {
	if l.key == "" {
		return ErrNoLock
	}
	if l.val == "" {
		return ErrNoData
	}
	err := l.UnlockString(ctx, l.key, l.val)
	if err != nil {
		return err
	}
	l.val = ""
	return nil
}

func (l *lock) Val() string {
	return l.val
}

func (l *lock) Key() string {
	return l.key
}

func (l *lock) LockString(ctx context.Context, key string, expire time.Duration) (string, error) {
	if key == "" {
		return "", errors.New("redis key cannot be empty")
	}
	val := uuid.NewString()
	err := l.redis.SetNX(ctx, key, val, expire).Err()
	return val, err
}

func (l *lock) UnlockString(ctx context.Context, key, val string) error {
	script := `
            if redis.call("get", KEYS[1]) == ARGV[1] then
                return redis.call("del", KEYS[1])
            else
                return 0
            end
`
	res, err := l.redis.Eval(ctx, script, []string{key}, val).Int()
	if err != nil {
		return err
	}
	if res == 0 {
		return nil
	}
	return nil
}
