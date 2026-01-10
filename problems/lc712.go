package problems

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	// dp[i][j] represents the minimum ASCII delete sum to make s1[0..i-1] and s2[0..j-1] equal
	dp := make([]int, n+1)

	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] + int(s2[j-1])
	}

	for i := 1; i <= m; i++ {
		// prev represents dp[i-1][j-1]
		prev := dp[0]
		// dp[0] represents dp[i][0]. Empty s2, need to delete all characters from s1
		dp[0] += int(s1[i-1])
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = prev
			} else {
				// Either delete s1[i-1] or s2[j-1]
				// if delete s1[i-1], dp[i][j] = dp[i-1][j] + ASCII(s1[i-1])
				// if delete s2[j-1], dp[i][j] = dp[i][j-1] + ASCII(s2[j-1])
				dp[j] = min(dp[j]+int(s1[i-1]), dp[j-1]+int(s2[j-1]))
			}
			prev = temp
		}
	}
	return dp[n]
}
