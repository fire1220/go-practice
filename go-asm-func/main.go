package main

import (
	"fmt"
	"go_asm_func/pkg"
)

func main() {
	pkg.Print("jock")
	pkg.Print2("hello world")
	fmt.Printf("pkg.Add:%#v\n", pkg.Add(1, 2))
	pkg.MyInfo()
}
