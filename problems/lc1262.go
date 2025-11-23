package problems

import "math"

func maxSumDivThree(nums []int) int {
	sum := 0
	min1, min11 := math.MaxInt, math.MaxInt
	min2, min22 := math.MaxInt, math.MaxInt
	for _, x := range nums {
		sum += x
		r := x % 3
		switch r {
		case 1:
			if x < min1 {
				min1, min11 = x, min1
			} else if x < min11 {
				min11 = x
			}
		case 2:
			if x < min2 {
				min2, min22 = x, min2
			} else if x < min22 {
				min22 = x
			}
		}
	}
	r := sum % 3
	if r == 0 {
		return sum
	}
	if r == 1 {
		x1 := min1
		x2 := math.MaxInt
		if min22 != math.MaxInt {
			x2 = min2 + min22
		}
		sub := min(x1, x2)
		if sub != math.MaxInt {
			return sum - sub
		}
		return 0
	}
	x1 := min2
	x2 := math.MaxInt
	if min11 != math.MaxInt {
		x2 = min1 + min11
	}
	sub := min(x1, x2)
	if sub != math.MaxInt {
		return sum - sub
	}
	return 0
}
