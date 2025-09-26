package problems

import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	// a < b < c
	for i := n - 1; i >= 2; i-- {
		c := nums[i]
		if nums[0]+nums[1] > c {
			res += (i + 1) * i * (i - 1) / 6
			break
		}
		if nums[i-2]+nums[i-1] <= nums[i] {
			continue
		}
		l, r := 0, i-1
		for l < r {
			a, b := nums[l], nums[r]
			if a+b > c {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}
	return res
}
