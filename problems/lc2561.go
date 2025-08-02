package problems

import "sort"

func minCost(basket1 []int, basket2 []int) int64 {
	N := len(basket1)
	cm := make(map[int]int, N*2)
	for _, x := range basket1 {
		cm[x]++
	}
	for _, x := range basket2 {
		cm[x]--
	}
	minv := basket1[0]
	cand := make([]int, 0, N)
	for k, v := range cm {
		minv = min(minv, k)
		if v < 0 {
			v = -v
		}
		if v%2 == 1 {
			return -1
		}
		v /= 2
		for i := 0; i < v; i++ {
			cand = append(cand, k)
		}
	}
	sort.Ints(cand)
	UB := len(cand) / 2
	bc := minv * 2
	cost := 0
	for i := range UB {
		cost += min(bc, cand[i])
	}
	return int64(cost)
}
