package ch4

import "fmt"

// Reverse a int slice
func Reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// RotateLeft will rotate a int slice
func Rotate(s []int, num int) []int {
	num = num % len(s)
	num += len(s)
	num %= len(s)
	for i := 0; i < len(s)-1; i++ {
		dst := (i + num) % len(s)
		s[i], s[dst] = s[dst], s[i]

		fmt.Println(i, dst)
	}

	return s
}
