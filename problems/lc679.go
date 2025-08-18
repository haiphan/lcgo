package problems

import "math"

const eps float64 = 0.00001

func getPossibleVal(a, b float64) []float64 {
	var opts []float64 = []float64{a + b, a - b, b - a, a * b}
	if b != 0.0 {
		opts = append(opts, a/b)
	}
	if a != 0.0 {
		opts = append(opts, b/a)
	}
	return opts
}
func judgeDfs(nums []float64) bool {
	if len(nums) == 1 {
		return math.Abs(nums[0]-24.0) < eps
	}
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			rest := make([]float64, 0, len(nums)-1)
			opts := getPossibleVal(a, b)
			for r := range nums {
				if r != i && r != j {
					rest = append(rest, nums[r])
				}
			}
			for _, x := range opts {
				rest = append(rest, x)
				if judgeDfs(rest) {
					return true
				}
				rest = rest[:len(rest)-1]
			}
		}
	}
	return false
}
func judgePoint24(cards []int) bool {
	nums := make([]float64, 4)
	for i, x := range cards {
		nums[i] = float64(x)
	}
	return judgeDfs(nums)
}
