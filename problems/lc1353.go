package problems

import "sort"

func heapifyBottom(heap []int) {
	i := len(heap) - 1

	for i > 0 {
		p := parent(i)
		if heap[p] <= heap[i] {
			return
		}
		heap[i], heap[p] = heap[p], heap[i]
		i = p
	}
}

func parent(i int) int {
	return ((i + 1) / 2) - 1
}

func maxEvents(events [][]int) int {
	N := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	cur := events[0][0]
	i, cnt := 0, 0
	eq := make([]int, 0, N)
	for i < N || len(eq) > 0 {
		for i < N && events[i][0] <= cur {
			eq = append(eq, events[i][1])
			heapifyBottom(eq)
			i++
		}
		for len(eq) > 0 && eq[0] < cur {
			eq = hPop(eq)
		}
		if len(eq) > 0 {
			eq = hPop(eq)
			cnt++
		}
		cur++
	}
	return cnt
}
