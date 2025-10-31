package problems

func maxProfit121(prices []int) int {
	ans := 0
	b := prices[0]
	for _, p := range prices {
		b = min(b, p)
		ans = max(ans, p-b)
	}
	return ans
}
