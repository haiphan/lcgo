package problems

import "sort"

func minCost(b1, b2 []int) int64 {
	n := len(b1)
	i, j := 0, 0
	count := make([]int, n)
	cost := make([]int, n)

	sort.Ints(b1)
	sort.Ints(b2)
	// swapPos is the current position in cost and count
	// swapCnt is the total number of groups that can be formed
	swapPos, swapCnt := 0, 0
	minCost := min(b1[0], b2[0]) << 1

	for i < n || j < n {
		v := 0
		cnt := 0
		if i < n && j < n && b1[i] == b2[j] {
			i++
			j++
			continue
		} else if i < n && (j == n || b1[i] < b2[j]) {
			v = b1[i]
			for i < n && b1[i] == v {
				cnt++
				i++
			}
		} else if j < n && (i == n || b2[j] < b1[i]) {
			v = b2[j]
			for j < n && b2[j] == v {
				cnt++
				j++
			}
		}
		if cnt&1 != 0 {
			return -1
		}
		cost[swapPos] = v
		count[swapPos] = cnt >> 1
		swapCnt += count[swapPos]
		swapPos++
	}

	if swapCnt&1 != 0 {
		return -1
	}
	swapCnt >>= 1
	swapPos = 0
	res := 0
	for swapCnt > 0 {
		used := min(count[swapPos], swapCnt)
		res += min(minCost, cost[swapPos]) * used
		swapCnt -= used
		swapPos++
	}
	return int64(res)
}
