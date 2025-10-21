package problems

import "sort"

func maxFrequencyI(nums []int, k int, numOperations int) int {
	L := len(nums)
	cm := make(map[int]int, len(nums))
	n := numOperations
	sort.Ints(nums)
	for _, x := range nums {
		cm[x]++
	}
	l, r, res := 0, 0, 0
	for _, v := range nums {
		for nums[l]+k < v {
			l++
		}
		for r+1 < L && v+k >= nums[r+1] {
			r++
		}
		cv := cm[v]
		cnt := r - l + 1
		change := min(n, cnt-cv)
		res = max(res, cv+change)
	}
	if res >= n {
		return res
	}
	l, r = 0, 0
	for ; r < L; r++ {
		for nums[l]+2*k < nums[r] {
			l++
		}
		res = max(res, min(n, r-l+1))
	}
	return res
}
