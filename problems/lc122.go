package problems

func maxProfit122(prices []int) int {
	N, p := len(prices), 0
	if N == 1 {
		return 0
	}
	for i := 1; i < N; i++ {
		p += max(0, prices[i]-prices[i-1])
	}
	return p
}
