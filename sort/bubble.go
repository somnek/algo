package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(l []int) []int {
	for i := 0; i < len(l)-1; i++ {
		for j := 0; j < len(l)-i-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
	return l
}

func randomNums(size int, max int) []int {
	a := []int{}
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(max))
	}
	return a
}

func main() {
	x := randomNums(10, 100)
	fmt.Println(bubbleSort(x))
}
