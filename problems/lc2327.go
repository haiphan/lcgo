package problems

func peopleAwareOfSecret(n int, delay int, forget int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	cur := 0
	for d := 2; d <= n; d++ {
		if d > delay {
			cur = (cur + dp[d-delay]) % MOD
		}
		if d > forget {
			cur = (cur - dp[d-forget] + MOD) % MOD
		}
		dp[d] = cur
	}
	cnt := 0
	for i := n - forget + 1; i <= n; i++ {
		cnt = (cnt + dp[i]) % MOD
	}
	return cnt
}
