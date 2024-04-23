package main

import (
	"context"
	"fmt"
	"go-unsafe/src"
	"reflect"
)

func main() {
	// user := people.People{Name: "jock", Nick: "fire"}
	// user.SetLike("篮球")
	// fmt.Println(user.GetLike())
	// src.GetPrivate(user, "like")

	a := context.Background()
	b := context.WithValue(a, "name", "jock")
	c := context.WithValue(b, "like", "足球")
	src.ContextKeys(c)
}

func TestStructByName() {
	type People struct {
		Name string
		Age  int
		Like []string
	}
	p := People{
		Name: "张三",
		Age:  12,
		Like: []string{"篮球", "排球"},
	}
	fmt.Println(StructByName(&p, "Like"))
	x := StructByName(&p, "Like")
	s, ok := x.([]string)
	fmt.Println(s, ok)
}

// StructByName 通过结构体的key名称获取值
func StructByName(data interface{}, key string) interface{} {
	v := reflect.ValueOf(data)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Type().Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	validateKey := false
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == key {
			validateKey = true
			break
		}
	}
	if validateKey {
		return v.FieldByName(key).Interface()
	}
	return nil
}
