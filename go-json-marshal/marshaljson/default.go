package marshaljson

import (
	"fmt"
	"reflect"
)

type defaultT struct {
}

func (d defaultT) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	fmt.Println("kkk")
	return reflect.ValueOf(defaultT{}), true
}
