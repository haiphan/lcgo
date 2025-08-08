package problems

var dp [201][201]float64

func calProb(a, b int) float64 {
	if a <= 0 && b > 0 {
		return 1.0
	}
	if a <= 0 && b <= 0 {
		return 0.5
	}
	if a > 0 && b <= 0 {
		return 0
	}
	if dp[a][b] != -1.0 {
		return dp[a][b]
	}
	dp[a][b] = 0.25 * (calProb(a-4, b) + calProb(a-3, b-1) + calProb(a-2, b-2) + calProb(a-1, b-3))
	return dp[a][b]
}

func soupServings(n int) float64 {
	if n > 5000 {
		return 1.0
	}
	for i := range 201 {
		for j := range 201 {
			dp[i][j] = -1.0
		}
	}
	n = (n + 24) / 25
	return calProb(n, n)
}
