package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type DateTime time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`"0000-00-00 00:00:00"`), nil
	}
	var stamp = fmt.Sprintf("%q", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func Marshal(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	if ref.Kind() != reflect.Struct {
		return nil, nil
	}
	var newTyp []reflect.StructField
	typ := ref.Type()
	typDate := reflect.TypeOf(DateTime{})
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		typeTemp := field.Type
		if typ.Field(i).Type.String() == "time.Time" {
			typeTemp = typDate
		}
		newTyp = append(newTyp, reflect.StructField{
			Name: field.Name,
			Type: typeTemp,
			Tag:  field.Tag,
		})
	}
	newStruct := reflect.New(reflect.StructOf(newTyp)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		field := newStruct.Field(i)
		oldField := ref.Field(i)
		if oldField.Type().String() != "time.Time" {
			field.Set(oldField)
			continue
		}
		if f, ok := oldField.Interface().(time.Time); ok {
			field.Set(reflect.ValueOf(DateTime(f)))
		}
	}
	return json.Marshal(newStruct.Interface())
}

type Good struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return Marshal(t)
}

func main() {
	good := Good{123, "jock", time.Now(), time.Now()}
	bytes, _ := json.Marshal(good)
	fmt.Printf("%s", bytes)
}
