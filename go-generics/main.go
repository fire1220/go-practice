package main

import (
	"fmt"
	"go-generics/dao"
	"go-generics/tools"
)

// go 泛型演示
func main() {
	fmt.Println("类展示：")
	d, err := dao.NewStudent().Generics().GetOne()
	fmt.Printf("%#v\t%#v\n", d, err) // model.Student{Id:2, Name:"李四", Like:"足球"}   <nil>
	x, err := dao.NewTeacher().Generics().GetOne()
	fmt.Printf("%#v\t%#v\n", x, err) // model.Teacher{Id:2, Name:"李四", Type:1}        <nil>

	fmt.Println("工具展示：")
	tool()
}

func tool() {
	arr := []int64{1, 2, 3, 4}
	x := tools.Array[int64](arr)
	fmt.Println(x.InArray(3))          // true
	fmt.Println(x.InArray(5))          // false
	fmt.Println(tools.InArray(3, arr)) // true
	fmt.Println(tools.InArray(5, arr)) // false
}
