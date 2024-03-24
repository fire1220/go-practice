package main

import (
	"fmt"
	"go_asm_func/pkg"
)

func main() {
	name := "jock"
	pkg.Print(name)
	pkg.Print2("hello world")
	fmt.Printf("pkg.Add:%#v\n", pkg.Add(1, 2))
}
