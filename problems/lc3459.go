package problems

var distCache [2000000]int
var dirs []int = []int{1, 1, -1, -1, 1}

func lenOfVDiagonal(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	L := 8 * m * n
	BASE := 8 * n
	for i := range L {
		distCache[i] = -1
	}
	isInbound := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < m
	}
	var dfs func(r, c, d, nxt, turn int) int
	dfs = func(r, c, d, nxt, turn int) int {
		k := turn + 2*d + 8*r + BASE*c
		if distCache[k] != -1 {
			return distCache[k]
		}
		v := 1
		nr, nc := r+dirs[d], c+dirs[d+1]
		if isInbound(nr, nc) && grid[nr][nc] == nxt {
			v = max(v, 1+dfs(nr, nc, d, 2-nxt, turn))
		}
		if turn == 0 {
			nd := (d + 1) % 4
			nr, nc := r+dirs[nd], c+dirs[nd+1]
			if isInbound(nr, nc) && grid[nr][nc] == nxt {
				v = max(v, 1+dfs(nr, nc, nd, 2-nxt, 1))
			}
		}
		distCache[k] = v
		return v
	}
	ans := 0
	for i, row := range grid {
		for j, cur := range row {
			if cur == 1 {
				for d := range 4 {
					ans = max(ans, dfs(i, j, d, 2, 0))
				}
			}
		}
	}
	return ans
}
