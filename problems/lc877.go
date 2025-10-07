package problems

func stoneGame(piles []int) bool {
	N := len(piles)
	dp := make([]int, N)
	copy(dp, piles)
	for d := 1; d < N; d++ {
		for i := 0; i+d < N; i++ {
			// take i or i+d
			dp[i] = max(piles[i]-dp[i+1], piles[i+d]-dp[i])
		}
	}
	return dp[0] > 0
}
