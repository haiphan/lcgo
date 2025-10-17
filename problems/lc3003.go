package problems

import "math/bits"

var CPRE, CSUF [10000]int

func cPack(parts, mask int) int {
	return parts<<32 | mask
}

func cUnpack(v int) (parts, mask int) {
	return v >> 32, v & 0xffffffff
}

func maxPartitionsAfterOperations(s string, k int) int {
	if k == 26 {
		return 1
	}
	var cc uint32
	for _, c := range s {
		cc |= 1 << (c - 'a')
	}
	if bits.OnesCount32(cc) < k {
		return 1
	}
	n := len(s)
	buildArr := func(arr *[10000]int, isPre bool) {
		start, end := 0, n-1
		step := 1
		if !isPre {
			start, end = end, start
			step = -1
		}
		mask, parts := 0, 0
		for i := start; i != end; i += step {
			charMask := 1 << (s[i] - 'a')
			mask |= charMask
			nc := bits.OnesCount32(uint32(mask))
			if nc > k {
				parts++
				mask = charMask
			}
			arr[i+step] = cPack(parts, mask)
		}
	}
	buildArr(&CPRE, true)
	buildArr(&CSUF, false)
	CSUF[n-1] = 0
	ans := 0
	for i := range n {
		pL, mL := cUnpack(CPRE[i])
		pR, mR := cUnpack(CSUF[i])
		cL, cR := bits.OnesCount32(uint32(mL)), bits.OnesCount32(uint32(mR))
		parts := pL + pR + 1
		count := bits.OnesCount32(uint32(mL | mR))
		if cL == k && cR == k && count < 26 {
			parts += 2
		} else if count >= k {
			parts += 1
		}
		ans = max(ans, parts)
	}
	return ans
}
