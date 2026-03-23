package problems

func numRollsToTarget(n int, k int, target int) int {
	M := 1000000007
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := min(target, i*k); j >= 0; j-- {
			mMax := min(j, k)
			dp[j] = 0
			for m := 1; m <= mMax; m++ {
				v := dp[j-m]
				dp[j] += v
			}
			dp[j] %= M
		}
	}
	return dp[target]
}
