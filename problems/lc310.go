package problems

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n, n)
	for i := range n {
		adj[i] = []int{}
	}
	deg := make([]int, n, n)
	done := make([]bool, n, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		deg[a]++
		deg[b]++
	}
	leaves := make([]int, 0, n)
	for i := range n {
		if deg[i] == 1 {
			leaves = append(leaves, i)
			done[i] = true
		}
	}
	remain := n
	for remain > 2 {
		remain -= len(leaves)
		nl := []int{}
		for _, l := range leaves {
			for _, next := range adj[l] {
				if done[next] {
					continue
				}
				deg[next]--
				if deg[next] == 1 {
					done[next] = true
					nl = append(nl, next)
				}
			}
		}
		leaves = nl
	}
	return leaves
}
