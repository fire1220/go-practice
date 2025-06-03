package main

import "fmt"

func main() {
	list := []int{4, 5, 6, 7, 8, 9}
	TreeHeap(list)
	fmt.Println(list) // [9 8 4 7 5 6]
}

func TreeHeap(list []int) {
	n := len(list)
	for i := n/2 - 1; i >= 0; i-- {
		down(list, i)
	}
}

func down(list []int, i int) {
	n := len(list)
	for {
		j := 2*i + 1 // left
		if j > n {
			break
		}
		if j < n-1 && list[j] < list[j+1] {
			j = j + 1 // right
		}
		if list[i] > list[j] {
			break
		}
		list[i], list[j] = list[j], list[i]
		i = j
	}
}
