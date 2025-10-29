package problems

import "math/bits"

func smallestNumber(n int) int {
	g := 1 << bits.Len(uint(n-1))
	if g == n {
		g <<= 1
	}
	return g - 1
}
