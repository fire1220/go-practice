package go_redis

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestHGetSet(t *testing.T) {
	ctx := context.Background()
	hkey := "user"
	key1 := "name"
	val1 := "jock1"
	key2 := "age"
	val2 := "11"
	fmt.Printf("%#v\n", new(RedisLock).Connect().HGet(ctx, hkey, key1).Val())
	fmt.Printf("%#v\n", new(RedisLock).Connect().HSet(ctx, hkey, key1, val1).Err())
	fmt.Printf("%#v\n", new(RedisLock).Connect().HGet(ctx, hkey, key1).Val()) // "jock1"
	fmt.Printf("%#v\n", new(RedisLock).Connect().HGet(ctx, hkey, key2).Val())
	fmt.Printf("%#v\n", new(RedisLock).Connect().HSet(ctx, hkey, key2, val2).Err())
	fmt.Printf("%#v\n", new(RedisLock).Connect().HGet(ctx, hkey, key2).String()) // "hget user age: 11"
	fmt.Printf("%#v\n", new(RedisLock).Connect().Del(ctx, hkey).Err())
}

func TestLock(t *testing.T) {
	ctx := context.Background()
	key := "test_123"
	val := "aaa"
	lock := new(RedisLock).Lock(ctx, key, val, 3*time.Second)
	fmt.Printf("%#v\n", lock)
	time.Sleep(2 * time.Second)
	a := new(RedisLock).Connect().Get(ctx, key).Val()
	fmt.Printf("%#v\n", a)
	unlock := new(RedisLock).Unlock(ctx, key, val)
	fmt.Printf("%#v\n", new(RedisLock).Connect().Get(ctx, key).Val())
	fmt.Printf("%#v\n", unlock)
}
