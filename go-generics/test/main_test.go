package test

import (
	"fmt"
	"go-generics/tools"
	"testing"
)

type MyInt int

func TestAdd(t *testing.T) {
	var a int8 = 1
	var b int8 = 2
	x := tools.Add(MyInt(a), MyInt(b))
	fmt.Println(x)
}
