package main

import _ "unsafe"

// 导出内部方法指向新名称(使用linkname必须导入unsafe包)
//go:linkname printstring runtime.printstring
func printstring()

// 注意：//go:linkname不能有空格
//go:linkname printnl runtime.printnl
func printnl()

// 自定义汇编函数
func myPrint(str string)

func main() {
	myPrint("hello")
}
