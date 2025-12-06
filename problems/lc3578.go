package problems

func countPartitions3578(nums []int, k int) int {
	n := len(nums)
	// dp[i]: number of ways to partition first i elements
	dp := make([]int, n+1)
	// prefix[i]: prefix sum of dp up to i
	prefix := make([]int, n+1)
	// minQ: increasing queue
	// maxQ: decreasing queue
	const MOD = 1e9 + 7
	minQ := make([]int, 0, n)
	maxQ := make([]int, 0, n)
	dp[0] = 1
	prefix[0] = 1
	// l pointer of the valid window
	l := 0
	for r, v := range nums {
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, r)
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, r)
		for len(maxQ) > 0 && len(minQ) > 0 && nums[maxQ[0]]-nums[minQ[0]] > k {
			if maxQ[0] == l {
				maxQ = maxQ[1:]
			}
			if minQ[0] == l {
				minQ = minQ[1:]
			}
			l++
		}
		dp[r+1] = prefix[r]
		if l > 0 {
			dp[r+1] = (dp[r+1] - prefix[l-1] + MOD) % MOD
		}
		prefix[r+1] = (prefix[r] + dp[r+1]) % MOD
	}

	return dp[n]
}
