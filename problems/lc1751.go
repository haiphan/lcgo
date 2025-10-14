package problems

import "sort"

func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	d1 := make([]int, n+1)
	d2 := make([]int, n+1)
	dp := [2][]int{d1, d2}
	next := make([]int, n)
	for j := range n {
		end := events[j][1]
		next[j] = sort.Search(n, func(i int) bool {
			return i > j && events[i][0] > end
		})
	}
	for j := 1; j <= k; j++ {
		cur := j & 1
		prev := cur ^ 1
		for i := n - 1; i >= 0; i-- {
			attend_i := events[i][2] + dp[prev][next[i]]
			skip_i := dp[cur][i+1]
			dp[cur][i] = max(attend_i, skip_i)
		}
	}
	return dp[k&1][0]
}
