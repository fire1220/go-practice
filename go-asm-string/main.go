package main

import (
	"fmt"
	"go-string/pkg"
)

func main() {
	fmt.Println()
	fmt.Printf("string pgk.Name:%#v\n", pkg.Name)
	fmt.Printf("string pgk.Name:%#v\n", pkg.Title)
	fmt.Println()
	fmt.Printf("slice pgk.Arr:%#v\n", pkg.Arr)
	fmt.Printf("slice pgk.Arr len:%#v\n", len(pkg.Arr))
	fmt.Printf("slice pgk.Arr Cap:%#v\n", cap(pkg.Arr))
	fmt.Println()
	fmt.Printf("string pgk.Age:%#v\n", pkg.Age)
	fmt.Println()
	fmt.Printf("string pgk.Id:%#v\n", pkg.Id)
	fmt.Printf("string pgk.UserId:%#v\n", pkg.UserId)
	fmt.Println()
}
