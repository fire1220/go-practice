package main

import (
	"fmt"
	"reflect"
	"time"
	_ "unsafe"
)

//go:linkname convT2E runtime.convT2E
func convT2E()

func getg() interface{}

func main() {
	go func() {
		id := GetGoId()
		fmt.Println(id)
		panic("test")
	}()
	time.Sleep(time.Second)
}

func GetGoId() int64 {
	g := getg()
	goId := reflect.ValueOf(g).FieldByName("goid").Int()
	return goId
}
