package problems

import "math/bits"

func kthCharacter(k int) byte {
	c := bits.OnesCount32(uint32(k - 1))
	return 'a' + byte(c)
}
