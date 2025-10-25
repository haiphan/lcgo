package problems

const WEEK_EARN int = 28

func totalMoney(n int) int {
	y := n % 7
	k := (n - y) / 7
	return k*WEEK_EARN + k*n - ((7 * k * (k + 1)) >> 1) + ((y * (y + 1)) >> 1)
}
