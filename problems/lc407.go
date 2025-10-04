package problems

type RCell struct {
	h, r, c int
}

func RLessThan(a, b RCell) bool {
	return a.h < b.h
}

var gridDirs [5]int = [5]int{0, 1, 0, -1, 0}

func RHPush(h []RCell, x RCell) []RCell {
	cur := len(h)
	h = append(h, x)
	for cur > 0 {
		p := (cur - 1) >> 1
		if RLessThan(h[cur], h[p]) {
			h[cur], h[p] = h[p], h[cur]
			cur = p
		} else {
			break
		}
	}
	return h
}

func RHPop(h []RCell) []RCell {
	last := h[len(h)-1]
	h[0] = last
	h = h[:len(h)-1]
	cur, l := 0, 1
	for l < len(h) {
		r := l + 1
		c := l
		if r < len(h) && RLessThan(h[r], h[l]) {
			c = r
		}
		if RLessThan(h[c], h[cur]) {
			h[c], h[cur] = h[cur], h[c]
			cur = c
			l = (c << 1) + 1
		} else {
			break
		}
	}
	return h
}

func trapRainWater(heightMap [][]int) int {
	hm := heightMap
	m, n := len(hm), len(hm[0])
	if m <= 2 || n <= 2 {
		return 0
	}
	pq := make([]RCell, 0, m*n)
	for i := range m {
		step := n - 1
		if i == 0 || i == m-1 {
			step = 1
		}
		for j := 0; j < n; j += step {
			pq = RHPush(pq, RCell{h: hm[i][j], r: i, c: j})
			hm[i][j] = -1
		}
	}
	ans, w := 0, 0
	for len(pq) > 0 {
		top := pq[0]
		pq = RHPop(pq)
		w = max(w, top.h)
		for i := range 4 {
			nr, nc := top.r+gridDirs[i], top.c+gridDirs[i+1]
			if nr < 0 || nr == m || nc < 0 || nc == n || hm[nr][nc] == -1 {
				continue
			}
			nh := hm[nr][nc]
			hm[nr][nc] = -1
			if nh < w {
				ans += w - nh
			}
			pq = RHPush(pq, RCell{h: nh, r: nr, c: nc})
		}
	}
	return ans
}
