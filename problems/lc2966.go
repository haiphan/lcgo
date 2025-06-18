package problems

import "sort"

func divideArray(nums []int, k int) [][]int {
	N := len(nums)
	M := N / 3
	sort.Ints(nums)
	res := make([][]int, 0, M)
	for i := 2; i < N; i += 3 {
		if nums[i]-nums[i-2] > k {
			return [][]int{}
		}
		res = append(res, []int{nums[i-2], nums[i-1], nums[i]})
	}
	return res
}
