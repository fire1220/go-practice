package test

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
	"time"
)

func TestForRange(t *testing.T) {
	// go1.22.0新增的写法
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("go1.22 has lift-off!")
}

var dataJson = `{"id":2,"name":"李四","like":"足球","type":1}`

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Like string `json:"like"`
}

func TestT1(t *testing.T) {
	var data Student
	err := jsoniter.UnmarshalFromString(dataJson, &data)
	fmt.Printf("%#v\n", data)
	fmt.Println(err)
}

func TestForGo(t *testing.T) {
	for v := range [...]int{1, 2, 3, 4, 5} {
		go func() {
			fmt.Printf("%v", v) // 在1.22.0之前是：55555；在之后则正常打印
		}()
	}
	fmt.Println()
	time.Sleep(3 * time.Second)
}
