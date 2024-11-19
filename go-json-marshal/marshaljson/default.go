package marshaljson

import (
	"reflect"
)

type defaultT struct {
}

func (d defaultT) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	return reflect.ValueOf(defaultT{}), true
}
