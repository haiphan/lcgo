package problems

func successfulPairs(spells []int, potions []int, success int64) []int {
	suc := int(success)
	maxp := 0
	for _, p := range potions {
		maxp = max(maxp, p)
	}
	dp := make([]int, maxp+1)
	for _, p := range potions {
		dp[p]++
	}
	for i := maxp - 1; i >= 0; i-- {
		dp[i] += dp[i+1]
	}
	res := make([]int, len(spells))
	for i, s := range spells {
		// s * p >= suc; p >= suc / s
		need := (suc + s - 1) / s
		if need <= maxp {
			res[i] = dp[need]
		}
	}
	return res
}
