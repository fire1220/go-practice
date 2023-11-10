package main

import (
	"fmt"
)

// 下面输出的结果是什么？为什么？

func main() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}

// 输出结果
// len(pollorder) =  5
// cap(pollorder) =  5
// len(lockorder) =  5
// cap(lockorder) =  5
