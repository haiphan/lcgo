package problems

import "sort"

func getPrevIndex(events [][]int, start int) int {
	l, r, res := 0, len(events)-1, -1
	for l <= r {
		m := (l + r) / 2
		if events[m][1] < start {
			res, l = m, m+1
		} else {
			r = m - 1
		}
	}
	return res
}
func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		e := events[i-1]
		prev := getPrevIndex(events, e[0])
		for j := 1; j <= k; j++ {
			skip := dp[i-1][j]
			take := e[2] + dp[prev+1][j-1]
			dp[i][j] = max(skip, take)
		}
	}

	return dp[n][k]
}
