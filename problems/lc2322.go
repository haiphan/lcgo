package problems

import "math"

func minimumScore(nums []int, edges [][]int) int {
	N := len(nums)
	adj := make([][]int, N)
	dp := make([]int, N)
	inId := make([]int, N)
	outId := make([]int, N)
	id := 0
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	var dfs func(prev, cur int)
	dfs = func(prev, cur int) {
		inId[cur] = id
		id++
		dp[cur] = nums[cur]
		for _, next := range adj[cur] {
			if next == prev {
				continue
			}
			dfs(cur, next)
			dp[cur] ^= dp[next]
		}
		outId[cur] = id
	}
	dfs(-1, 0)
	t0 := dp[0]
	res := math.MaxInt
	for i := 1; i < N; i++ {
		ti := dp[i]
		for j := i + 1; j < N; j++ {
			tj := dp[j]
			a, b, c := ti, tj, t0^ti^tj
			if inId[j] > inId[i] && inId[j] < outId[i] {
				a, b, c = tj, ti^tj, t0^ti
			} else if inId[i] > inId[j] && inId[i] < outId[j] {
				a, b, c = ti, tj^ti, t0^tj
			}
			res = min(res, max(a, b, c)-min(a, b, c))
		}
	}
	return res
}
