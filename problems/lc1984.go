package problems

import "sort"

func minimumDifference1984(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	h, l := nums[k-1], nums[0]
	ans := h - l
	for i := k; i < n; i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}
