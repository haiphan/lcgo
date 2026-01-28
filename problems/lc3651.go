package problems

import (
	"slices"
)

func minCost3651(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if k > 0 && grid[0][0] > grid[m-1][n-1] {
		return 0
	}
	BIG := 1600000
	maxv := 0
	for _, row := range grid {
		maxv = max(maxv, slices.Max(row))
	}
	// sufMinF[x] = min cost to reach any cell with value >= x
	sufMinF := make([]int, maxv+2)
	for i := range sufMinF {
		sufMinF[i] = BIG
	}
	// minF[x] = min cost to reach any cell with value == x in current iteration
	minF := make([]int, maxv+1)
	// dp[j + 1] = min cost to reach cell in current row and column j
	dp := make([]int, n+1)
	// iterate k + 1 times because we can increase cell values up to k times
	for range k + 1 {
		for i := range minF {
			minF[i] = BIG
		}
		// reset f. By default all inf except f[1]
		for i := range dp {
			dp[i] = BIG
		}
		// starting point. We pay negative cost to enter the cell
		dp[1] = -grid[0][0]
		for _, row := range grid {
			for j, x := range row {
				// left f[j] + x, up f[j+1] + x, and from cells with value >= x
				dp[j+1] = min(dp[j]+x, dp[j+1]+x, sufMinF[x])
				// update minF for this cell value
				minF[x] = min(minF[x], dp[j+1])
			}
		}

		done := true
		for i := maxv; i >= 0; i-- {
			// update sufMinF
			minDist := min(sufMinF[i+1], minF[i])
			if minDist < sufMinF[i] {
				sufMinF[i] = minDist
				done = false
			}
		}
		// if no update in sufMinF, we can stop early
		if done {
			break
		}
	}
	return dp[n]
}
