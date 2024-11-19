package marshaljson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func MarshalFormat(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()
	newField := make([]reflect.StructField, 0, ref.NumField())
	isNeedNewStruct := false
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		for _, tabName := range tabList {
			if field.Tag.Get(tabName) == "" {
				continue
			}
			tm, ok := tabMap[tabName]
			if !ok {
				continue
			}
			if tm.restrain != "" && field.Type.String() != tm.restrain {
				continue
			}
			fieldType = tm.refTypOf
			isNeedNewStruct = true
			break
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}
	if !isNeedNewStruct {
		return json.Marshal(p)
	}

	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldField := ref.Field(i)
		oldTyp := typ.Field(i)
		var newFieldVal reflect.Value
		newFieldVal = oldField
		for _, tabName := range tabList {
			if oldTyp.Tag.Get(tabName) == "" {
				continue
			}
			tm, ok := tabMap[tabName]
			if !ok {
				continue
			}
			if oldTyp.Type.String() != tm.restrain {
				continue
			}
			if tm.fun == nil {
				continue
			}
			newVal, ok := tm.fun.typeConv(oldField, oldTyp)
			if ok {
				newFieldVal = newVal
			}
			break
		}

		fmt.Println(newStruct.Field(i).Type().Name())
		newStruct.Field(i).Set(newFieldVal)
	}
	return json.Marshal(newStruct.Interface())
}

func MarshalFormat1(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()

	newField := make([]reflect.StructField, 0, ref.NumField())
	dateTimeReflectType := reflect.TypeOf(dateTime{})
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		if field.Type.String() == "time.Time" {
			fieldType = dateTimeReflectType
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}

	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldField := ref.Field(i)
		oldFieldType := typ.Field(i)
		if oldField.Type().String() != "time.Time" {
			newStruct.Field(i).Set(oldField)
			continue
		}
		if v, ok := oldField.Interface().(time.Time); ok {
			newStruct.Field(i).Set(reflect.ValueOf(dateTime{t: v, tag: oldFieldType.Tag}))
		}
	}

	return json.Marshal(newStruct.Interface())
}
