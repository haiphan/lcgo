package problems

func maxArea(height []int) int {
	N := len(height)
	l, r := 0, N-1
	res := 0
	for l < r {
		res = max(res, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
