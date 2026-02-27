package problems

import "math"

func minOperations3666(s string, k int) int {
	n := len(s)
	nz := 0
	for _, c := range s {
		if c == '0' {
			nz++
		}
	}
	if n == k {
		if nz == 0 {
			return 0
		}
		if nz == n {
			return 1
		}
		return -1
	}
	ans := math.MaxInt
	if nz%2 == 0 {
		m := max((nz+k-1)/k, (nz+n-k-1)/(n-k))
		m += m & 1
		ans = min(ans, m)
	}

	if (nz % 2) == (k % 2) {
		m := max((nz+k-1)/k, (n-nz+n-k-1)/(n-k))
		m += 1 - (m & 1)
		ans = min(ans, m)
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
