package ch4

import "crypto/sha256"

func popCount(x [32]byte) int {
	n := 0
	for _, val := range x {
		x := val
		for x != 0 {
			x = x & (x - 1)
			n++
		}
	}

	return n
}

func SHA256Diff(a, b []byte) int {
	c1 := sha256.Sum256(a)
	c2 := sha256.Sum256(b)

	n1 := popCount(c1)
	n2 := popCount(c2)

	return n1 - n2
}
