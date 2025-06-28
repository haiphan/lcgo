package problems

import "sort"

func maxSubsequence(nums []int, k int) []int {
	N := len(nums)
	iArr := make([]int, N)
	for i := range N {
		iArr[i] = i
	}
	sort.Slice(iArr, func(i, j int) bool {
		return nums[iArr[j]] < nums[iArr[i]]
	})
	iArr = iArr[:k]
	sort.Ints(iArr)
	for i, j := range iArr {
		iArr[i] = nums[j]
	}
	return iArr
}
