package problems

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k-1+maxPts {
		return 1
	}
	mp := float64(maxPts)
	dp := make([]float64, n+1)
	dp[0] = 1.0
	res, wsum := 0.0, 1.0
	for i := 1; i <= n; i++ {
		dp[i] = wsum / mp
		if i < k {
			wsum += dp[i]
		} else {
			res += dp[i]
		}
		if i >= maxPts {
			wsum -= dp[i-maxPts]
		}
	}
	return res
}
