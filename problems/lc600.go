package problems

import "math/bits"

var FVAL [30]int

func fillFib() {
	if FVAL[0] == 1 {
		return
	}
	FVAL[0], FVAL[1] = 1, 2
	for i := 2; i < len(FVAL); i++ {
		FVAL[i] = FVAL[i-1] + FVAL[i-2]
	}
}

func findIntegers(n int) int {
	fillFib()
	L := bits.Len64(uint64(n))
	prev := 0
	ans := 0
	for i := L; i >= 0; i-- {
		if ((n >> i) & 1) == 1 {
			ans += FVAL[i]
			if prev == 1 {
				return ans
			}
			prev = 1
		} else {
			prev = 0
		}
	}
	return ans + 1
}
