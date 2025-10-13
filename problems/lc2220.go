package problems

import "math/bits"

func minBitFlips(start int, goal int) int {
	return bits.OnesCount32(uint32(start ^ goal))
}
