package problems

import "sort"

func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	dp := make([]int, n+1)
	ndp := make([]int, n+1)
	// next[i]: the next event index that can be attended after attending event i
	next := make([]int, n)
	for j := range n {
		end := events[j][1]
		next[j] = sort.Search(n, func(i int) bool {
			return i > j && events[i][0] > end
		})
	}
	for j := 1; j <= k; j++ {
		for i := n - 1; i >= 0; i-- {
			attend_i := events[i][2] + dp[next[i]]
			skip_i := ndp[i+1]
			ndp[i] = max(attend_i, skip_i)
		}
		dp, ndp = ndp, dp
	}
	return dp[0]
}
