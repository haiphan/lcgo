package problems

func trap(height []int) int {
	h := height
	N := len(height)
	lMax, rMax := 0, 0
	l, r := 0, N-1
	ans := 0
	for l <= r {
		if h[l] <= h[r] {
			if h[l] > lMax {
				lMax = h[l]
			} else {
				ans += lMax - h[l]
			}
			l++
		} else {
			if h[r] > rMax {
				rMax = h[r]
			} else {
				ans += rMax - h[r]
			}
			r--
		}
	}
	return ans
}
