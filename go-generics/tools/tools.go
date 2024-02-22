package tools

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
