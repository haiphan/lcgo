package problems

func count(cur, n int) int {
	res, nb := 0, cur+1
	for cur <= n {
		res += min(n+1, nb) - cur
		cur *= 10
		nb *= 10
	}
	return res
}
func findKthNumber(n int, k int) int {
	cur, i := 1, 1
	for i < k {
		steps := count(cur, n)
		if i+steps <= k {
			i += steps
			cur++
		} else {
			cur *= 10
			i++
		}
	}
	return cur
}
