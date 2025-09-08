package problems

func getNoZeroIntegers(n int) []int {
	a, b := 0, 0
	t := 1
	c := 0
	for ; n > 0; n, t, c = n/10-c, t*10, 0 {
		r := n % 10
		if r == 0 {
			d := 5 * t
			a += d
			b += d
			c = 1
		} else if r == 1 && n > 10 {
			a += 6 * t
			b += 5 * t
			c = 1
		} else {
			da := r / 2
			a += da * t
			b += (r - da) * t
		}
	}
	return []int{a, b}
}
