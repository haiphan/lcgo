package problems

func stoneGameII(piles []int) int {
	n := len(piles)
	base := n + 1
	suf := make([]int, n)
	dp := make([]int, base*base)
	suf[n-1] = piles[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = piles[i] + suf[i+1]
	}
	for i := n - 1; i >= 0; i-- {
		for m := 1; m <= n; m++ {
			k := i + base*m
			if i+2*m >= n {
				dp[k] = suf[i]
				continue
			}
			for x := 1; x <= 2*m; x++ {
				dp[k] = max(dp[k], suf[i]-dp[i+x+base*max(x, m)])
			}
		}
	}

	return dp[base]
}
