package problems

func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	r := 0
	for i := 1; i <= n; i++ {
		p := intPow(i, x)
		if p > n {
			break
		}
		r += p
		for j := min(n, r); j >= p; j-- {
			dp[j] += dp[j-p]
			if dp[j] > MOD {
				dp[j] -= MOD
			}
		}
	}
	return dp[n]
}
