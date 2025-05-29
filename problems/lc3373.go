package problems

func assignParity(cur, prev, p int, adj [][]int, par []int) {
	par[cur] = p
	for _, next := range adj[cur] {
		if next == prev {
			continue
		}
		assignParity(next, cur, 1-p, adj, par)
	}
}

func maxTargetNodes2(edges1 [][]int, edges2 [][]int) []int {
	N, M := len(edges1)+1, len(edges2)+1
	adj1, adj2 := getAdj(edges1), getAdj(edges2)
	par1, par2 := make([]int, N), make([]int, M)
	res := make([]int, N)
	assignParity(0, -1, 0, adj1, par1)
	assignParity(0, -1, 0, adj2, par2)
	odd2, odd1 := 0, 0
	for _, x := range par2 {
		odd2 += x
	}
	MP2 := max(odd2, M-odd2)
	for _, x := range par1 {
		odd1 += x
	}
	even1 := N - odd1
	for i := range N {
		cnt := even1
		if par1[i] == 1 {
			cnt = odd1
		}
		res[i] = MP2 + cnt
	}
	return res
}
