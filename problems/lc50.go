package problems

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var res float64 = 1
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x = x * x
		n = n >> 1
	}
	if neg {
		return float64(1) / res
	}
	return res
}
