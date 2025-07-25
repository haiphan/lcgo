package problems

import "math"

func increasingTriplet(nums []int) bool {
	// p2, p1, cur
	p2, p1 := math.MaxInt32, math.MaxInt32
	for _, x := range nums {
		if x <= p2 {
			p2 = x
		} else if x <= p1 {
			p1 = x
		} else {
			return true
		}
	}
	return false
}
