package main

import (
	"fmt"
)

func maxVal(l []int) (max int) {
	for i := 0; i < len(l)-1; i++ {
		if l[i] > l[i+1] {
			l[i], l[i+1] = l[i+1], l[i]
		}
	}
	return l[len(l)-1]
}

func main() {
	l := []int{2, 3, 5, 0, 4, 1}
	fmt.Println(maxVal(l))

}
