package problems

func minDeletionSize960(strs []string) int {
	// n := len(strs)
	m := len(strs[0])
	// dp[i] = max number of sorted columns ending with column i
	dp := make([]int, m)
	dp[0] = 1
	// max columns to keep
	keep := 1

	for i := 1; i < m; i++ {
		dp[i] = 1
		for j := range i {
			canAdd := true
			for _, s := range strs {
				if s[j] > s[i] {
					canAdd = false
					break
				}
			}

			if canAdd {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		keep = max(keep, dp[i])
	}
	return m - keep
}
