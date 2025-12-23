package problems

import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	suf := make([]int, n)
	suf[n-1] = events[n-1][2]
	for i := n - 2; i >= 0; i-- {
		suf[i] = max(suf[i+1], events[i][2])
	}
	ans := 0
	for i, e := range events {
		cur := e[2]
		j := -1
		l, r := i+1, n-1
		for l <= r {
			m := (l + r) >> 1
			if events[m][0] > e[1] {
				j = m
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if j != -1 {
			cur += suf[j]
		}
		ans = max(ans, cur)
	}
	return ans
}
