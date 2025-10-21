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
		for r < L && v+k >= nums[r] {
			r++
		}
		cnt := r - l
		if cnt <= res {
			continue
		}
		cv := cm[v]
		change := min(n, cnt-cv)
		res = max(res, cv+change)
	}
	if res >= n {
		return res
	}
	tk := k << 1
	l, r = 0, 0
	for ; r < L && (l+res) < L; r++ {
		for r-l+1 > n || nums[l]+tk < nums[r] {
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
