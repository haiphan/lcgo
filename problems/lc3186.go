package problems

import "sort"

func maximumTotalDamage(power []int) int64 {
	cm := make(map[int]int, len(power))
	for _, x := range power {
		cm[x]++
	}
	n := len(cm)
	values := make([]int, n)
	vi := 0
	for k := range cm {
		values[vi] = k
		vi++
	}
	sort.Ints(values)
	dp := make([]int, n)
	dp[0] = values[0] * cm[values[0]]
	for i := 1; i < n; i++ {
		v := values[i]
		take := v * cm[values[i]]
		j := i - 1
		for j >= 0 && values[j]+2 >= v {
			j--
		}
		if j >= 0 {
			take += dp[j]
		}
		dp[i] = max(dp[i-1], take)
	}
	return int64(dp[n-1])
}
