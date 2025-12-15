package problems

func clumsy(n int) int {
	if n <= 2 {
		return n
	}
	if n <= 4 {
		return n + 3
	}
	pick := [4]int{1, 2, 2, -1}
	return n + pick[(n-4)%4]
}
