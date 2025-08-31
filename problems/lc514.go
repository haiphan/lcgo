package problems

import "math"

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	base := max(m, n)
	dp := make([]int, base*base)
	cp := make([][]int, 26)
	for i := range ring {
		code := int(ring[i] - 'a')
		cp[code] = append(cp[code], i)
	}
	var dfs func(i, j int) int // i pos, j key idx
	dfs = func(i, j int) int {
		if j == n {
			return 0
		}
		k := i*base + j
		if dp[k] > 0 {
			return dp[k]
		}
		res := math.MaxInt
		code := int(key[j] - 'a')
		for _, next := range cp[code] {
			d := abs(next - i)
			d = min(d, m-d)
			res = min(res, d+dfs(next, j+1))
		}
		res++
		dp[k] = res
		return res
	}

	return dfs(0, 0)
}
