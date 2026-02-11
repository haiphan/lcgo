package problems

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		a, b := indexes[i], indexes[j]
		return startTime[a] < startTime[b] || (startTime[a] == startTime[b] && endTime[a] < endTime[b])
	})
	//[start, maxProfit]
	done := make([][2]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		Index := indexes[i]
		start, end, p := startTime[Index], endTime[Index], profit[Index]
		if len(done) == 0 {
			done = append(done, [2]int{start, p})
			continue
		}
		if end <= done[0][0] {
			prev := sort.Search(len(done), func(j int) bool {
				return done[j][0] < end
			}) - 1
			// prev--
			p += done[prev][1]
		}
		if p > done[len(done)-1][1] {
			if start == done[len(done)-1][0] {
				done[len(done)-1][1] = p
			} else {
				done = append(done, [2]int{start, p})
			}
		}
	}
	return done[len(done)-1][1]
}
