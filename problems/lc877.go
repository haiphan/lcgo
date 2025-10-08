package problems

func stoneGame(piles []int) bool {
	N := len(piles)
	dp := make([]int, N)
	copy(dp, piles)
	for d := 1; d < N; d++ {
		for i := 0; i+d < N; i++ {
			// dp[i] is the max score difference can get from piles[i..i+d]
			// if take i, opponent can get dp[i+1], so net is piles[i]-dp[i+1]
			// if take i+d, opponent can get dp[i], so net is piles[i+d]-dp[i]
			dp[i] = max(piles[i]-dp[i+1], piles[i+d]-dp[i])
		}
	}
	return dp[0] > 0
}
