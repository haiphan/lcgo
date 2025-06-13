package problems

import "sort"

func isValid(nums []int, p, u, n int) bool {
	i, cnt := 0, 0
	L1 := n - 1
	for i < L1 {
		if nums[i+1]-nums[i] <= u {
			i += 2
			cnt++
			if cnt == p {
				return true
			}
		} else {
			i++
		}
	}
	return false
}
func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	N := len(nums)
	sort.Ints(nums)
	l, r := 0, nums[N-1]-nums[0]
	ans := r
	for l <= r {
		m := (l + r) / 2
		if isValid(nums, p, m, N) {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
