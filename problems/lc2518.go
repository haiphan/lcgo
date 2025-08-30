package problems

func countPartitions(nums []int, k int) int {
	nSum := 0
	for _, x := range nums {
		nSum += x
	}
	if nSum < 2*k {
		return 0
	}
	n := len(nums)
	dp := make([]int, k)
	dp[0] = 1
	for _, x := range nums {
		for i := k - 1; i >= x; i-- {
			dp[i] = (dp[i] + dp[i-x]) % MOD
		}
	}
	bad := 0
	for _, x := range dp {
		bad = (bad + x) % MOD
	}
	bad = (bad * 2) % MOD
	total := 1
	for i := 0; i < n; i++ {
		total = (total << 1) % MOD
	}
	return (total + MOD - bad) % MOD
}
