package problems

const WEEK_EARN int = 28

func totalMoney(n int) int {
	if n <= 7 {
		return (n * (n + 1)) >> 1
	}
	n -= 7
	return WEEK_EARN + n + totalMoney(n)
}
