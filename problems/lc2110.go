package problems

func getDescentPeriods(prices []int) int64 {
	N := len(prices)
	cnt, res := 1, 1
	for i := 1; i < N; i++ {
		if prices[i] == prices[i-1]-1 {
			cnt++
		} else {
			cnt = 1
		}
		res += cnt
	}
	return int64(res)
}
