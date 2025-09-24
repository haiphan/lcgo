package problems

import (
	"math"
)

func calDiff(a, b, k int, s string, pre []int) int {
	N := len(s)
	maxDiff := math.MinInt32
	minVal := [4]int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}
	l, r := 0, k-1
	pa, pb := 0, 0
	for ; r < N; r++ {
		ca := pre[a*(N+1)+r+1]
		cb := pre[b*(N+1)+r+1]
		for r-l+1 >= k && cb-pb >= 2 {
			parity := ((pa & 1) << 1) + (pb & 1)
			minVal[parity] = min(minVal[parity], pa-pb)
			l++
			pa = pre[a*(N+1)+l]
			pb = pre[b*(N+1)+l]
		}
		// for a, diffenrent parity. for b the same
		maxDiff = max(maxDiff, ca-cb-minVal[(((ca&1)^1)<<1)+(cb&1)])
	}
	return maxDiff
}

func maxDifferenceII(s string, k int) int {
	N := len(s)
	pre := make([]int, 5*(N+1)) // pre[a*5+i] = count of digit a before position i
	for i := 0; i < N; i++ {
		d := int(s[i] - '0')
		for j := 0; j < 5; j++ {
			pre[j*(N+1)+i+1] = pre[j*(N+1)+i]
		}
		pre[d*(N+1)+i+1]++
	}
	maxDiff := math.MinInt32
	for a := range 5 {
		if pre[a*(N+1)+N] == 0 {
			continue
		}
		for b := range 5 {
			if a == b || pre[b*(N+1)+N] == 0 {
				continue
			}
			maxDiff = max(maxDiff, calDiff(a, b, k, s, pre))
		}
	}
	if maxDiff == math.MinInt32 {
		return 0
	}
	return maxDiff
}
