package problems

import (
	"sort"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type Element struct {
	p int
	c int
}

func byPriority(a, b interface{}) int {
	priorityA := a.(Element).p
	priorityB := b.(Element).p
	return -utils.IntComparator(priorityA, priorityB)
}
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	N := len(profits)
	items := make([]Element, N)
	total := 0
	maxC := capital[0]
	for i := 0; i < N; i++ {
		total += profits[i]
		maxC = max(maxC, capital[i])
		items[i] = Element{c: capital[i], p: profits[i]}
	}
	if maxC <= w && k >= N {
		return total + w
	}
	itemQ := pq.NewWith(byPriority)
	sort.Slice(items, func(i, j int) bool {
		return items[i].c < items[j].c
	})
	i := 0
	for k > 0 {
		for i < N && items[i].c <= w {
			itemQ.Enqueue(items[i])
			i++
		}
		if itemQ.Size() == 0 {
			return w
		}
		cur, _ := itemQ.Dequeue()
		w += cur.(Element).p
		k--
	}
	return w
}
