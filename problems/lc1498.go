package problems

import "sort"

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	N := len(nums)
	pc := make([]int, N)
	pc[0] = 1
	for i := 1; i < N; i++ {
		pc[i] = pc[i-1] * 2 % MOD
	}
	r := N - 1
	res := 0
	for l, v := range nums {
		if l > r || v > target {
			break
		}
		for l <= r && v+nums[r] > target {
			r--
		}
		if l <= r {
			res = (res + pc[r-l]) % MOD
		}
	}
	return res
}
