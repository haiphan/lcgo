package problems

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adj := make([][]int, n)
	ans := 0
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	var dfs func(prev, cur int) int
	dfs = func(prev, cur int) int {
		sum := values[cur] % k
		for _, nxt := range adj[cur] {
			if nxt == prev {
				continue
			}
			sum += dfs(cur, nxt)
			sum %= k
		}
		if sum == 0 {
			ans++
		}
		return sum
	}
	dfs(-1, 0)
	return ans
}
