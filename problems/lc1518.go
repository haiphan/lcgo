package problems

func numWaterBottles(numBottles int, numExchange int) int {
	cnt := 0
	n, k := numBottles, numExchange
	for n >= k {
		n += 1 - k
		cnt += k
	}
	return cnt + n
}
