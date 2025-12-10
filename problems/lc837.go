package problems

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k-1+maxPts {
		return 1
	}
	mp := float64(maxPts)
	// dp[i]: probability to get points i
	dp := make([]float64, n+1)
	dp[0] = 1.0
	// wsum: window sum
	res, wsum := 0.0, 1.0
	for i := 1; i <= n; i++ {
		dp[i] = wsum / mp
		if i < k {
			// still drawing
			wsum += dp[i]
		} else {
			// stop drawing, result accumulate
			res += dp[i]
		}
		if i >= maxPts {
			// shrink the window
			wsum -= dp[i-maxPts]
		}
	}
	return res
}
