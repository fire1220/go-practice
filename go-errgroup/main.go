package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	// 下面goErrGroup和goWithGroup两者是等价的
	goErrGroup()            // 用errgroup同步原语启动协成
	goWithGroup()           // 用等待组的方式启动协成
	goErrGroupWithContext() // 通过上下文关闭功能退出
}

// 用errgroup同步原语启动协成
func goErrGroup() {
	g := errgroup.Group{}
	for k := range []int{1, 2, 3, 4, 5} {
		k := k
		g.Go(func() error {
			fmt.Println(k)
			return fmt.Errorf("失败：%v\n", k)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}

// 用等待组的方式启动协成
func goWithGroup() {
	wg := sync.WaitGroup{}
	list := []int{1, 2, 3, 4, 5}
	chErr := make(chan error, len(list))
	for k := range list {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			fmt.Println(k)
			chErr <- fmt.Errorf("失败：%v\n", k)
		}(k)
	}
	wg.Wait()
	if d, ok := <-chErr; ok {
		fmt.Println(d)
	} else {
		fmt.Println("ok")
	}
}

// errgroup关闭上下文功能
func goErrGroupWithContext() {
	g, ctx := errgroup.WithContext(context.Background())
	for k := range []int{1, 2, 3, 4, 5} {
		k := k
		g.Go(func() error {
			fmt.Println(k)
			return fmt.Errorf("失败：%v\n", k)
		})
		if isCancel(ctx) {
			break
		}
		time.Sleep(time.Second)
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}

func isCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}
	return false
}
