package problems

import "sort"

func getMaxGap(bars []int) int {
	sort.Ints(bars)
	cur := 1
	mg := 1
	for i, x := range bars {
		if i > 0 && x == bars[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		mg = max(mg, cur)
	}
	return mg + 1
}

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	gap := min(getMaxGap(hBars), getMaxGap(vBars))
	return gap * gap
}
