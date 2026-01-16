package problems

import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hh, vv := hFences, vFences
	hh = append(hh, 1)
	hh = append(hh, m)
	vv = append(vv, 1)
	vv = append(vv, n)
	sort.Ints(hh)
	sort.Ints(vv)
	if len(hh) > len(vv) {
		hh, vv = vv, hh
		m, n = n, m
	}
	hgap := make(map[int]bool, min(m, 600))
	for i, a := range hh {
		for j := i + 1; j < len(hh); j++ {
			g := hh[j] - a
			hgap[g] = true
		}
	}
	maxl := 0
	vl := len(vv)
	for i, a := range vv {
		maxPossibleGap := n - a
		if maxPossibleGap <= maxl {
			break
		}
		for j := vl - 1; j > i; j-- {
			g := vv[j] - a
			if g <= maxl {
				break
			}
			if hgap[g] {
				maxl = g
				break
			}
		}
	}
	if maxl == 0 {
		return -1
	}
	return (maxl * maxl) % MOD
}
