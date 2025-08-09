package problems

import "math/bits"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return bits.OnesCount32(uint32(n)) == 1
}
