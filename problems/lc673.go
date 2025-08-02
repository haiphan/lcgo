package problems

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	count := make([]int, n)
	maxF, maxCnt := 0, 0
	for i := range n {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				curF := 1 + dp[j]
				if curF > dp[i] {
					dp[i] = 1 + dp[j]
					count[i] = count[j]
				} else if curF == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxF {
			maxF = dp[i]
			maxCnt = count[i]
		} else if dp[i] == maxF {
			maxCnt += count[i]
		}
	}
	return maxCnt
}
