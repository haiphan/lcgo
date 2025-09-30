package problems

import "sort"

func findLongestChain(pairs [][]int) int {
	N := len(pairs)
	if N == 1 {
		return 1
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	prev, cnt := -1001, 0
	for _, p := range pairs {
		if prev < p[0] {
			prev = p[1]
			cnt++
		}
	}
	return cnt
}
