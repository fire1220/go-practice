package main

import "fmt"

func Xadd(ptr *uint32, new uint32) uint32 // 注释：原子操作，交换位置并且返回，把ptr指针里的值和new的相互交换后返回ptr指针对应的值

func main() {
	var p1 uint32 = 12
	var p2 uint32 = 13
	fmt.Println(p1, p2) // 12 13
	old := Xadd(&p1, p2)
	fmt.Println(old, p1) // 25 25
}
