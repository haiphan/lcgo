package problems

import "sort"

func maxTaxiEarnings(n int, rides [][]int) int64 {
	m, ri := len(rides), 0
	dp := make([]int, n+1)
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] < rides[j][0]
	})
	for i := 1; i <= n; i++ {
		dp[i] = max(dp[i], dp[i-1])
		for ri < m && rides[ri][0] == i {
			cur := rides[ri]
			e, t := cur[1], cur[2]
			dp[e] = max(dp[e], dp[i]+e-i+t)
			ri++
		}
	}
	return int64(dp[n])
}
