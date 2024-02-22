package dao

import (
	"fmt"
	"go-generics/model"
)

type teacher struct {
}

func NewTeacher() *teacher {
	return new(teacher)
}
func (t *teacher) Generics() *Generics[model.Teacher] {
	return new(Generics[model.Teacher])
}

func (t *teacher) Study() {
	fmt.Println("study")
}
