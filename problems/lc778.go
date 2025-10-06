package problems

type ECell struct {
	e, r, c int
}

func ELessThan(a, b ECell) bool {
	return a.e < b.e
}

func EHPush(h []ECell, x ECell) []ECell {
	cur := len(h)
	h = append(h, x)
	for cur > 0 {
		p := (cur - 1) >> 1
		if ELessThan(h[cur], h[p]) {
			h[cur], h[p] = h[p], h[cur]
			cur = p
		} else {
			break
		}
	}
	return h
}

func EHPop(h []ECell) []ECell {
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]
	cur, l := 0, 1
	for l < len(h) {
		r := l + 1
		c := l
		if r < len(h) && ELessThan(h[r], h[l]) {
			c = r
		}
		if ELessThan(h[c], h[cur]) {
			h[c], h[cur] = h[cur], h[c]
			cur = c
			l = (c << 1) + 1
		} else {
			break
		}
	}
	return h
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	pq := make([]ECell, 0, n*n)
	t := 0
	pq = EHPush(pq, ECell{e: grid[0][0], r: 0, c: 0})
	grid[0][0] = -1
	for len(pq) > 0 {
		top := pq[0]
		pq = EHPop(pq)
		t = max(t, top.e)
		if top.r == n-1 && top.c == n-1 {
			return t
		}
		for i := range 4 {
			nr, nc := top.r+gridDirs[i], top.c+gridDirs[i+1]
			if nr < 0 || nr == n || nc < 0 || nc == n || grid[nr][nc] == -1 {
				continue
			}
			ne := grid[nr][nc]
			grid[nr][nc] = -1
			pq = EHPush(pq, ECell{e: ne, r: nr, c: nc})
		}
	}
	return t
}
