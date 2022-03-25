package main

import (
	"fmt"
)

func minVal(l []int) int {
	minIdx := 0
	for i := 0; i < len(l)-1; i++ {
		if l[i] < l[minIdx] {
			minIdx = i
		}
	}
	return l[minIdx]
}

func main() {
	l := []int{-1, 4, 3, 1, 2, 0, 5}
	fmt.Println(minVal(l))
}
