package problems

import "math/bits"

func pyramidTransition(bottom string, allowed []string) bool {
	var bToH [64]int
	n := len(bottom)
	for _, a := range allowed {
		cl, cr, ct := a[0]&31, a[1]&31, a[2]&31
		bToH[cl<<3+cr] |= (1 << ct)
	}
	pyramid := make([]int, n)
	for i, c := range bottom {
		pyramid[n-1] |= int(c&31) << (i * 3)
	}
	seen := make([]bool, 1<<((n-1)*3))
	var backtrack func(row, col int) bool
	backtrack = func(row, col int) bool {
		if row < 0 {
			return true
		}
		rowVal := pyramid[row]
		if seen[rowVal] {
			return false
		}
		if col == row+1 {
			seen[rowVal] = true
			return backtrack(row-1, 0)
		}
		prevVal := pyramid[row+1]
		cl, cr := (prevVal>>(col*3))&7, (prevVal>>((col+1)*3))&7
		chars := uint32(bToH[cl<<3+cr])
		for chars != 0 {
			ct := bits.TrailingZeros32(chars)
			pyramid[row] &^= 7 << (col * 3)
			pyramid[row] |= ct << (col * 3)
			if backtrack(row, col+1) {
				return true
			}
			chars &= (chars - 1)
		}
		return false
	}
	return backtrack(n-2, 0)
}
