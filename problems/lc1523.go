package problems

func countOdds(low int, high int) int {
	n := high - low + 1
	if n%2 == 0 {
		return n >> 1
	}
	return (n >> 1) + (low % 2)
}
