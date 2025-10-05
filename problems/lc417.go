package problems

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pp := make([][]bool, m)
	aa := make([][]bool, m)
	for i := range m {
		aa[i] = make([]bool, n)
		pp[i] = make([]bool, n)
	}
	res := make([][]int, 0)
	var dfs func(r, c, prev int, visit [][]bool)
	dfs = func(r, c, prev int, visit [][]bool) {
		if r < 0 || r == m || c < 0 || c == n || heights[r][c] < prev || visit[r][c] {
			return
		}
		visit[r][c] = true
		h := heights[r][c]
		dfs(r-1, c, h, visit)
		dfs(r+1, c, h, visit)
		dfs(r, c-1, h, visit)
		dfs(r, c+1, h, visit)
	}
	for c := 0; c < n; c++ {
		dfs(0, c, -1, pp)
		dfs(m-1, c, -1, aa)
	}
	for r := 0; r < m; r++ {
		dfs(r, 0, -1, pp)
		dfs(r, n-1, -1, aa)
	}
	for r := range m {
		for c := range n {
			if aa[r][c] && pp[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}
