package problems

import (
	"math"
)

func maxStability(n int, edges [][]int, k int) int {
	orig_uf := NewUnionFind(n)
	min_str := math.MaxInt32
	max_str := 0
	for i := range edges {
		max_str = max(max_str, edges[i][2])
		if edges[i][3] == 0 {
			continue
		}
		min_str = min(min_str, edges[i][2])
		if !orig_uf.Union(edges[i][0], edges[i][1]) {
			return -1
		}
	}
	check := func(sta int) bool {
		if sta > min_str {
			return false
		}
		uf := NewUnionFind(n)
		for i := range n {
			uf.rank[i] = orig_uf.rank[i]
			uf.par[i] = orig_uf.par[i]
		}
		upgrade := make([][2]int, 0, n)
		for i := range edges {
			if edges[i][3] == 1 {
				continue
			}
			if edges[i][2] >= sta {
				uf.Union(edges[i][0], edges[i][1])
			} else if edges[i][2]*2 >= sta {
				upgrade = append(upgrade, [2]int{edges[i][0], edges[i][1]})
			}
		}
		r := k
		for i := range upgrade {
			u, v := upgrade[i][0], upgrade[i][1]
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
