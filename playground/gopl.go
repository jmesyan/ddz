package main

import (
	"fmt"
)

//
// 1
// 1 2
// 1 2 3
// 1 2 3 4
// 1 2 3 4 5

func reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// TODO: ROTATE

func main() {
	vec := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}

	for _, v := range vec {
		reverse(v)
		fmt.Println(v)
	}
	fmt.Println("Hello world")
}
