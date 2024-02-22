package tools

import "golang.org/x/exp/constraints"

func InArray[T comparable](needle T, haystack []T) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

type Array[T constraints.Ordered] []T

func (a Array[T]) InArray(needle T) bool {
	for _, val := range a {
		if val == needle {
			return true
		}
	}
	return false
}

type intSet interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Add[T intSet](a, b T) T {
	return a + b
}
