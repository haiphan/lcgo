package problems

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	// tn = tn1 + tn2 + tn3
	tn, tn1, tn2, tn3 := 0, 1, 1, 0
	for i := 3; i <= n; i++ {
		tn = tn1 + tn2 + tn3
		tn1, tn2, tn3 = tn, tn1, tn2
	}
	return tn
}
