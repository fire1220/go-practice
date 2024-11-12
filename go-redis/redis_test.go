package go_redis

import (
	"context"
	"fmt"
	"testing"
)

func TestHGetSet(t *testing.T) {
	ctx := context.Background()
	hkey := "user"
	key1 := "name"
	val1 := "jock1"
	key2 := "age"
	val2 := "11"
	fmt.Printf("%#v\n", Connection().HGet(ctx, hkey, key1).Val())
	fmt.Printf("%#v\n", Connection().HSet(ctx, hkey, key1, val1).Err())
	fmt.Printf("%#v\n", Connection().HGet(ctx, hkey, key1).Val()) // "jock1"
	fmt.Printf("%#v\n", Connection().HGet(ctx, hkey, key2).Val())
	fmt.Printf("%#v\n", Connection().HSet(ctx, hkey, key2, val2).Err())
	fmt.Printf("%#v\n", Connection().HGet(ctx, hkey, key2).String()) // "hget user age: 11"
	fmt.Printf("%#v\n", Connection().Del(ctx, hkey).Err())
}
