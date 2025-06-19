package problems

import "sort"

func partitionArray(nums []int, k int) int {
	N := len(nums)
	sort.Ints(nums)
	cnt, minV := 1, nums[0]
	for i := 1; i < N; i++ {
		if nums[i]-minV > k {
			cnt++
			minV = nums[i]
		}
	}
	return cnt
}
