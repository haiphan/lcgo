package problems

func getAdj(edges [][]int) [][]int {
	N := len(edges) + 1
	adj := make([][]int, N)
	for i := range N {
		adj[i] = make([]int, 0)
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	return adj
}

func dfs(adj [][]int, cur, prev, k int) int {
	if k < 0 {
		return 0
	}
	res := 1
	for _, next := range adj[cur] {
		if next == prev {
			continue
		}
		res += dfs(adj, next, cur, k-1)
	}
	return res
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	N, M := len(edges1)+1, len(edges2)+1
	adj1, adj2 := getAdj(edges1), getAdj(edges2)
	ML2 := 0
	for i := range M {
		ML2 = max(ML2, dfs(adj2, i, -1, k-1))
	}
	res := make([]int, N)
	for i := range N {
		res[i] = ML2 + dfs(adj1, i, -1, k)
	}
	return res
}
