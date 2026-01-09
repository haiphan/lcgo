package problems

import "slices"

type NumPack struct {
	old, computed, i int
}

func decodeNum(mapping []int, x int) int {
	if x == 0 {
		return mapping[0]
	}
	t := 1
	y := 0
	for x > 0 {
		d := x % 10
		y += t * mapping[d]
		t *= 10
		x /= 10
	}
	return y
}
func sortJumbled(mapping []int, nums []int) []int {
	arr := make([]NumPack, len(nums))
	for i, x := range nums {
		arr[i] = NumPack{old: x, computed: decodeNum(mapping, x), i: i}
	}
	slices.SortFunc(arr, func(a, b NumPack) int {
		if a.computed == b.computed {
			return a.i - b.i
		}
		return a.computed - b.computed
	})
	for i := range arr {
		nums[i] = arr[i].old
	}
	return nums
}
