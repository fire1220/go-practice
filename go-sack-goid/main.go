package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	go func() {
		fmt.Println(GetGoId())
		panic("验证") // 验证协成ID（panic的时候会打印出协成ID）
	}()
	time.Sleep(1 * time.Second)
}

// GetGoId 获取Go协成ID
func GetGoId() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id : %v", err))
	}
	return int64(id)
}
