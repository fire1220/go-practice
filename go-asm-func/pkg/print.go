package pkg

import (
	"fmt"
	_ "unsafe"
)

// 注意：//go:linkname不能有空格
//
//go:linkname printnl runtime.printnl
func printnl()

// 导出内部方法指向新名称(使用linkname必须导入unsafe包)
//
//go:linkname printstring runtime.printstring
func printstring()

func printStringData(str string) {
	fmt.Println(str)
}

func Print(str string)

func Print2(str string)
