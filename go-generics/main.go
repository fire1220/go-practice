package main

import (
	"fmt"
	"go-generics/dao"
	"go-generics/tools"
)

// go 泛型演示
func main() {
	fmt.Println("类展示：")
	d, err := dao.NewStudent().Gen().GetOne()
	// model.Student{Id:2, Name:"李四", Like:"足球"}   <nil>
	fmt.Printf("%#v\t%#v\n", d, err)
	d, err = dao.NewStudent().GetOne()
	// model.Student{Id:2, Name:"李四", Like:"足球"}   <nil>
	fmt.Printf("%#v\t%#v\n", d, err)

	x, err := dao.NewTeacher().Gen().GetOne()
	// model.Teacher{Id:2, Name:"李四", Type:1}        <nil>
	fmt.Printf("%#v\t%#v\n", x, err)
	x, err = dao.NewTeacher().GetOne()
	// model.Teacher{Id:2, Name:"李四", Type:1}        <nil>
	fmt.Printf("%#v\t%#v\n", x, err)
}

func tool() {
	arr := []int64{1, 2, 3, 4}
	x := tools.Array[int64](arr)
	fmt.Println(x.InArray(3))          // true
	fmt.Println(x.InArray(5))          // false
	fmt.Println(tools.InArray(3, arr)) // true
	fmt.Println(tools.InArray(5, arr)) // false
}
