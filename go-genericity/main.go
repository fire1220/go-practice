package main

import "fmt"

func InArray[T comparable](needle T, haystack []T) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

type Array[T int | int8 | int16 | int64] []T

func (a Array[T]) InArray(needle T) bool {
	for _, val := range a {
		if val == needle {
			return true
		}
	}
	return false
}

func main() {
	arr := []int64{1, 2, 3, 4}
	x := Array[int64](arr)
	fmt.Println(x.InArray(3))    // true
	fmt.Println(x.InArray(5))    // false
	fmt.Println(InArray(3, arr)) // true
	fmt.Println(InArray(5, arr)) // false
}
