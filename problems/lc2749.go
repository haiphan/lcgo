package problems

import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
	if num2 == 0 {
		return bits.OnesCount64(uint64(num1))
	}
	for k := 1; k <= 60; k++ {
		x := num1 - k*num2
		if x < k || x <= 0 {
			continue
		}
		if k >= bits.OnesCount64(uint64(x)) {
			return k
		}
	}
	return -1
}
