package main

import (
	"fmt"
	"time"
	"unsafe"
)

// type g struct的152字节位置是goid字段(目前go1.9;go1.10;go1.16是152，go1.8是192)
const gGoIdOffset = 152

func getg() uint

func main() {
	go func() {
		fmt.Println(GetGoroutineId())
		panic("test")
	}()
	time.Sleep(time.Second)
}

func GetGoroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + gGoIdOffset))
	return *p
}
