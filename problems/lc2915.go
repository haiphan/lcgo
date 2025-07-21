package problems

func lengthOfLongestSubsequence(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = -1
	}
	for _, n := range nums {
		for t := target; t >= n; t-- {
			prev := t - n
			if dp[prev] == -1 {
				continue
			}
			dp[t] = max(dp[t], 1+dp[prev])
		}
	}

	return dp[target]
}
