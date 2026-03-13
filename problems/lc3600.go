package problems

import (
	"math"
	"sort"
)

func maxStability(n int, edges [][]int, k int) int {
	orig_uf := NewUnionFind(n)
	min_str := math.MaxInt32
	max_str := 0
	optional := make([][3]int, 0, len(edges))
	for i := range edges {
		max_str = max(max_str, edges[i][2])
		if edges[i][3] == 0 {
			optional = append(optional, [3]int{edges[i][0], edges[i][1], edges[i][2]})
			continue
		}
		min_str = min(min_str, edges[i][2])
		if !orig_uf.Union(edges[i][0], edges[i][1]) {
			return -1
		}
	}
	sort.Slice(optional, func(i, j int) bool {
		return optional[i][2] > optional[j][2]
	})

	check := func(sta int) bool {
		if sta > min_str {
			return false
		}
		uf := NewUnionFind(n)
		for i := range n {
			uf.rank[i] = orig_uf.rank[i]
			uf.par[i] = orig_uf.par[i]
		}

		// Free optional edges first: w >= sta.
		idx := 0
		for idx < len(optional) && optional[idx][2] >= sta {
			uf.Union(optional[idx][0], optional[idx][1])
			idx++
		}

		r := k
		// Then upgrade-eligible optional edges: ceil(sta/2) <= w < sta.
		for idx < len(optional) && optional[idx][2]*2 >= sta {
			u, v := optional[idx][0], optional[idx][1]
			idx++
			if uf.Find(u) == uf.Find(v) {
				continue
			}
			if r == 0 {
				return false
			}
			uf.Union(u, v)
			r--
		}
		for i := range n {
			if uf.Find(i) != uf.Find(0) {
				return false
			}
		}
		return true
	}
	l, r := -1, max_str<<1
	for l < r {
		// add 1 to avoid infinite loop when l and r are adjacent
		m := l + (r-l+1)>>1
		if check(m) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
