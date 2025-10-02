package problems

func maxBottlesDrunk(numBottles int, numExchange int) int {
	n, k := numBottles, numExchange
	cnt := 0
	for n >= k {
		cnt += k
		n += 1 - k
		k++
	}
	return cnt + n
}
