package problems

type SCell struct {
	d, r, c int
}

func SCLessThan(a, b SCell) bool {
	return a.d > b.d
}

func SHPush(h []SCell, x SCell) []SCell {
	cur := len(h)
	h = append(h, x)
	for cur > 0 {
		p := (cur - 1) >> 1
		if SCLessThan(h[cur], h[p]) {
			h[cur], h[p] = h[p], h[cur]
			cur = p
		} else {
			break
		}
	}
	return h
}

func SHPop(h []SCell) []SCell {
	last := h[len(h)-1]
	h[0] = last
	h = h[:len(h)-1]
	cur, l := 0, 1
	for l < len(h) {
		r := l + 1
		c := l
		if r < len(h) && SCLessThan(h[r], h[l]) {
			c = r
		}
		if SCLessThan(h[c], h[cur]) {
			h[c], h[cur] = h[cur], h[c]
			cur = c
			l = (c << 1) + 1
		} else {
			break
		}
	}
	return h
}

func getSG(grid [][]int) [][]int {
	n := len(grid)
	q := make([][2]int, 0, n*n)
	for i := range n {
		for j := range n {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
			}
		}
	}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		r, c := top[0], top[1]
		d := grid[r][c]
		for i := range 4 {
			nr, nc := r+drs[i], c+dcs[i]
			if nr < 0 || nr == n || nc < 0 || nc == n || grid[nr][nc] != 0 {
				continue
			}
			grid[nr][nc] = d + 1
			q = append(q, [2]int{nr, nc})
		}
	}
	for i := range n {
		for j := range n {
			grid[i][j]--
		}
	}
	return grid
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}
	grid = getSG(grid)
	pq := make([]SCell, 0, n*n)
	pq = SHPush(pq, SCell{d: grid[0][0], r: 0, c: 0})
	visit := make([]bool, n*n)
	visit[0] = true
	for len(pq) > 0 {
		top := pq[0]
		pq = SHPop(pq)
		d, r, c := top.d, top.r, top.c
		if r == n-1 && c == n-1 {
			return d
		}
		if d == 0 {
			return 0
		}
		for i := range 4 {
			nr, nc := top.r+drs[i], top.c+dcs[i]
			k := n*nr + nc
			if nr < 0 || nr == n || nc < 0 || nc == n || visit[k] {
				continue
			}
			visit[k] = true
			pq = SHPush(pq, SCell{d: min(d, grid[nr][nc]), r: nr, c: nc})
		}
	}

	return 0
}
