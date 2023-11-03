package main

import (
	"fmt"
	_ "unsafe"
)

// 导出内部方法指向新名称(使用linkname必须导入unsafe包)
//go:linkname printstring runtime.printstring
func printstring(s string)

//go:linkname printnl runtime.printnl
func printnl()

var helloWorld = "hello world"

func main()

// 当前包下的汇编可以直接调用
func printString(str string) {
	fmt.Println(str)
}
