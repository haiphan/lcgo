package problems

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	L := len(nums)
	cm := make([]int, L)
	n := numOperations
	sort.Ints(nums)
	l, r := 0, 0
	for r < L {
		for r < L && nums[r] == nums[l] {
			r++
		}
		cnt := r - l
		for l < r {
			cm[l] = cnt
			l++
		}
	}
	l, r = 0, 0
	res := 0
	for i, v := range nums {
		for nums[l]+k < v {
			l++
		}
		for r < L && v+k >= nums[r] {
			r++
		}
		cnt := r - l
		if cnt <= res {
			continue
		}
		cv := cm[i]
		change := min(n, cnt-cv)
		res = max(res, cv+change)
	}
	if res >= n {
		return res
	}
	l, r = 0, 0
	tk := (k << 1)
	for ; r < L; r++ {
		for (r-l+1) > n || nums[l]+tk < nums[r] {
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
