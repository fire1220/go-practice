package go_redis

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	ctx := context.Background()
	key := "test_123"
	obj := NewLock()
	err := obj.Lock(ctx, key, 3*time.Second)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	// time.Sleep(4 * time.Second)

	fmt.Printf("%#v\n", Connection().Get(ctx, key).Val())
	err = obj.Unlock(ctx)
	fmt.Printf("%#v\n", Connection().Get(ctx, key).Val())
	fmt.Printf("%#v\n", err)
	err = obj.Unlock(ctx)
	fmt.Printf("%#v\n", err)
}
