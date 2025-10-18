package problems

import "sort"

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	minv := nums[0] - k - 1
	cnt := 0
	for _, x := range nums {
		if x+k > minv {
			minv = max(minv+1, x-k)
			cnt++
		}
	}
	return cnt
}
