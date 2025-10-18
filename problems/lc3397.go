package problems

import "sort"

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	minv := nums[0] - k - 1
	cnt := 0
	for _, x := range nums {
		nxt := min(max(minv+1, x-k), x+k)
		if nxt > minv {
			minv = nxt
			cnt++
		}
	}
	return cnt
}
