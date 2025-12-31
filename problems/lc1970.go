package problems

func latestDayToCross(row int, col int, cells [][]int) int {
	MD := [5]int{0, 1, 0, -1, 0}
	n := row * col
	top, bot := n, n+1
	par := make([]int, n+2)
	seen := make([]bool, n)
	for i := range n + 2 {
		par[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != par[x] {
			par[x] = find(par[x])
		}
		return par[x]
	}
	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			par[px] = py
		}
	}
	for d := n - 1; d >= 0; d-- {
		r, c := cells[d][0]-1, cells[d][1]-1
		id := r*col + c
		seen[id] = true
		if r == 0 {
			union(id, top)
		} else if r == row-1 {
			union(id, bot)
		}
		for i := range 4 {
			nr, nc := r+MD[i], c+MD[i+1]
			nid := nr*col + nc
			if nr >= 0 && nr < row && nc >= 0 && nc < col && seen[nid] {
				union(id, nid)
			}
		}
		// remove this cell makes top and bot connected
		// so the last day to cross is d
		if find(top) == find(bot) {
			return d
		}
	}
	return 0
}
