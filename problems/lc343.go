package problems

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	r := n % 3
	c3 := n / 3
	if r == 1 {
		r = 4
		c3--
	} else if r == 0 {
		r = 1
	}
	return intPow(3, c3) * r
}
