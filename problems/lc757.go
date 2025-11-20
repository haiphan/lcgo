package problems

import "sort"

const RIGHTP int = 1e8 + 1

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] > intervals[j][0]
	})
	p1, p2 := RIGHTP, RIGHTP
	ans := 0
	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		need := 2
		if p1 <= r {
			need--
		}
		if p2 <= r {
			need--
		}
		if need == 0 {
			continue
		}
		if need == 1 {
			ans++
			if p1 == l {
				p2 = l + 1
			} else {
				p1, p2 = l, p1
			}
		} else {
			ans += 2
			p1, p2 = l, l+1
		}
	}
	return ans
}
