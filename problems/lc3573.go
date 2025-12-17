package problems

func maximumProfit(prices []int, k int) int64 {
	x0 := prices[0]
	n := len(prices)
	dp := make([]int, n)
	res := 0
	for t := 1; t <= k; t++ {
		buy, sell := -x0, x0
		prev := dp[t-1]
		for i := t; i < n; i++ {
			old := dp[i]
			dp[i] = max(dp[i-1], buy+prices[i], sell-prices[i])
			res = max(res, dp[i])
			buy = max(buy, prev-prices[i])
			sell = max(sell, prev+prices[i])
			prev = old
		}
	}

	return int64(res)
}
