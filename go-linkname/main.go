package main

import (
	"fmt"
	_ "unsafe"
)

// 随机数
//go:linkname FastRand runtime.fastrand
func FastRand() uint32

// 时间戳
//go:linkname nanotime1 runtime.nanotime1
func nanotime1() int64

func main() {
	fmt.Println(FastRand())
	fmt.Println(nanotime1())
}
