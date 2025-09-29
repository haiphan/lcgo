package problems

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	if n == 3 {
		return values[0] * values[1] * values[2]
	}
	// dp[i][j] = min score polygon [i..j]
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	for d := 2; d < n; d++ {
		for i := 0; i+d < n; i++ {
			j := i + d
			minScore := math.MaxInt
			e := values[i] * values[j]
			for k := i + 1; k < j; k++ {
				minScore = min(minScore, e*values[k]+dp[i][k]+dp[k][j])
			}
			dp[i][j] = minScore
		}
	}
	return dp[0][n-1]
}
