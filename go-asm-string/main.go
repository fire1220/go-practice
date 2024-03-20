package main

import (
	"go-string/pkg"
)

func main() {
	println(pkg.Name)
	println(pkg.Title)
	// str := pkg.Name
	// t := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	// var ret []byte
	// sTemp := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	// sTemp.Len = t.Len
	// sTemp.Cap = t.Len
	// sTemp.Data = t.Data
	// fmt.Println(ret)
	// ret[0] = 111
	// fmt.Println(ret)

	// str := pkg.Name
	// ret := []byte(str)
	// ret[0] = 111
	// fmt.Println(ret)
}
