package problems

func stoneGame(piles []int) bool {
	N := len(piles)
	dp := make([][]int, N)
	for i := N - 1; i >= 0; i-- {
		dp[i] = make([]int, N)
		dp[i][i] = piles[i]
		for j := i + 1; j < N; j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][N-1] > 0
}
