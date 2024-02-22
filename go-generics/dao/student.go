package dao

import (
	"go-generics/model"
)

type student struct {
	Name string
}

func NewStudent() *student {
	return new(student)
}

func (s *student) Generics() *Generics[model.Student] {
	return new(Generics[model.Student])
}

func (s *student) One() (model.Student, error) {
	return model.Student{Id: 1, Name: "法外狂徒-张三", Like: "篮球"}, nil
}
