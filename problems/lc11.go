package problems

func maxArea(height []int) int {
	N := len(height)
	l, r := 0, N-1
	res := 0
	for l < r {
		v := 0
		if height[l] <= height[r] {
			v = height[l] * (r - l)
			l++
		} else {
			v = height[r] * (r - l)
			r--
		}
		res = max(res, v)
	}
	return res
}
