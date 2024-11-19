package marshaljson

import (
	"reflect"
)

type defaultT struct {
	tag reflect.StructTag
}

func (d defaultT) MarshalJSON() ([]byte, error) {
	format := d.tag.Get(tabDefault)

	return []byte(`"` + format + `"`), nil
}

func (d defaultT) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	return reflect.ValueOf(defaultT{tag: typ.Tag}), true
}
