package main

import (
	"fmt"
	"go-generics/dao"
)

// go 泛型演示
func main() {
	d, err := dao.NewStudent().Generics().GetOne()
	fmt.Printf("%#v\t%#v\n", d, err)
	x, err := dao.NewTeacher().Generics().GetOne()
	fmt.Printf("%#v\t%#v\n", x, err)
}
