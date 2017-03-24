package ch2

var pc []byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// PopCount calculate how many non-zero bits are there in x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountByLoop is another PopCount, but run in loop
func PopCountByLoop(x uint64) int {
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

// PopCountByShifting is another PopCount, calculate non-zero bits by shifting them
func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}

	return n
}

// PopCountByClearing is another PopCount, x & (x - 1) will clean the last non-zero bit in x
func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
