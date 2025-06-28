package problems

const MAXIV int = 2147483647

func reverse(x int) int {
	neg := x < 0
	if neg {
		x = -x
	}
	y := 0
	for x > 0 {
		d := x % 10
		if y > ((MAXIV - d) / 10) {
			return 0
		}
		y = 10*y + d
		x /= 10
	}
	if neg {
		return -y
	}
	return y
}
