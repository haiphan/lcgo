package problems

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	N, res := len(startTime), 0
	k1 := k - 1
	N1 := N - 1
	pref := 0
	for i := range N {
		pref += endTime[i] - startTime[i]
		if i > k1 {
			pref -= endTime[i-k] - startTime[i-k]
		}
		start := 0
		if i >= k {
			start = endTime[i-k]
		}
		end := eventTime
		if i < N1 {
			end = startTime[i+1]
		}
		res = max(res, end-start-pref)
	}
	return res
}
