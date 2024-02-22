package dao

import (
	jsoniter "github.com/json-iterator/go"
)

type Generics[T any] struct {
}

var dataJson = `{"id":2,"name":"李四","like":"足球","type":1}`

func (p Generics[T]) GetOne() (T, error) {
	var data T
	err := jsoniter.UnmarshalFromString(dataJson, &data)
	return data, err
}
