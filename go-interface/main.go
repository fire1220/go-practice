package main

import "fmt"

type IFace interface {
	Add(a, b int) int
}

type F float32

func (f F) Add(x, y int) int {
	return x + y
}

func main() {
	var a, b float32 = 1, 2
	fmt.Println(a) // eface
	test(F(b))     // iface
}

func test(s IFace) {
	println(s.Add(10, 12))
}
