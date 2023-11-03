package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
	_ "unsafe"
)

//go:linkname convT2E runtime.convT2E
func convT2E()

// func getg() uint
func getg() uint

// func getGoroutine() interface{}
func getGoroutine() interface{}

func main() {
	go func() {
		id := GetGoroutineId()
		fmt.Println(id)
		panic("test")
	}()
	time.Sleep(time.Second)
}

func GetGoroutineId() int64 {
	g := getg()
	var gGoIdOffset uintptr
	// 通过类型获取偏移量（效率比较高）
	if f, ok := reflect.TypeOf(getGoroutine()).FieldByName("goid"); ok {
		gGoIdOffset = f.Offset
	} else {
		panic("can not find g.goId field")
	}
	p := (*int64)(unsafe.Pointer(uintptr(g) + gGoIdOffset))
	return *p
}
