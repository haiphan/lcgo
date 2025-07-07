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
		digits = append([]int{c}, digits...)
	}
	return digits
}
