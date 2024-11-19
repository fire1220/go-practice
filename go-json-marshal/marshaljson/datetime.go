package marshaljson

import (
	"reflect"
	"strings"
	"time"
)

type dateTime struct {
	t   time.Time
	tag reflect.StructTag
}

func (d dateTime) MarshalJSON() ([]byte, error) {
	t := d.t
	formatData := d.tag.Get("datetime")
	format, ok := strings.CutSuffix(formatData, "omitempty")
	format = strings.Trim(format, ",")
	if format == "" {
		format = time.DateTime
	}
	mapTime := map[string]string{
		time.DateTime: "0000-00-00 00:00:00",
		time.DateOnly: "0000-00-00",
		time.TimeOnly: "00:00:00",
	}
	if t.IsZero() {
		if ok {
			return []byte(`""`), nil
		}
		if v, ok := mapTime[format]; ok {
			return []byte(`"` + v + `"`), nil
		} else {
			return []byte(`""`), nil
		}
	}
	return []byte(`"` + t.Format(format) + `"`), nil
}

func (d dateTime) typeConv(field reflect.Value, typ reflect.StructField) (reflect.Value, bool) {
	if v, ok := field.Interface().(time.Time); ok {
		return reflect.ValueOf(dateTime{t: v, tag: typ.Tag}), true
	}
	return reflect.Value{}, false
}
