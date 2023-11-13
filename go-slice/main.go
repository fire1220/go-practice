package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	test3()
}

// 下面输出的结果是什么？为什么？
func test1() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
	// 输出结果
	// len(pollorder) =  5
	// cap(pollorder) =  5
	// len(lockorder) =  5
	// cap(lockorder) =  5
}

// 下面是否会panic？为什么？
func test2() {
	world := `world`
	baseStr := `hello ` + world
	tempStr := *(*reflect.StringHeader)(unsafe.Pointer(&baseStr))
	var s []byte
	tempSlice := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	tempSlice.Cap = tempStr.Len
	tempSlice.Len = tempStr.Len
	tempSlice.Data = tempStr.Data
	fmt.Println(s)
	s[1] = 111
	fmt.Println(s)
	fmt.Println(*(*string)(unsafe.Pointer(&s)))

	// 不会panic，因为经过字符串拼接后，底层采用slice存储，并且通过指针转换成string的，所以最终存储的还是slice，汇编代码没有rodata只读标记
}

// 下面是否会panic？为什么？
func test3() {
	baseStr := `hello world`
	tempStr := *(*reflect.StringHeader)(unsafe.Pointer(&baseStr))
	var s []byte
	tempSlice := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	tempSlice.Cap = tempStr.Len
	tempSlice.Len = tempStr.Len
	tempSlice.Data = tempStr.Data
	s[1] = 111
	fmt.Println(s)
	// 会panic，因为string在汇编代码里添加了rodata只读标记
}
