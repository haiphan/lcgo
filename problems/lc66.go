package problems

func plusOne(digits []int) []int {
	c := 1
	N := len(digits)
	for i := N - 1; i >= 0; i-- {
		digits[i] += c
		c = 0
		if digits[i] > 9 {
			c = digits[i] / 10
			digits[i] %= 10
		}
	}
	if c > 0 {
		result := make([]int, N+1)
		result[0] = c
		copy(result[1:], digits)
		digits = result
	}
	return digits
}
