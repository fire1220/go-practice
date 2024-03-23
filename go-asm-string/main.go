package main

import (
	"fmt"
	"go-string/pkg"
)

func main() {
	fmt.Println("asm string:")
	fmt.Printf("string pgk.Name:%#v\n", pkg.Name)
	fmt.Printf("string pgk.Name:%#v\n", pkg.Title)
	fmt.Println()
	fmt.Println("asm slice:")
	fmt.Printf("slice pgk.Arr:%#v\n", pkg.Arr)
	fmt.Printf("slice pgk.Arr len:%#v\n", len(pkg.Arr))
	fmt.Printf("slice pgk.Arr Cap:%#v\n", cap(pkg.Arr))
}
