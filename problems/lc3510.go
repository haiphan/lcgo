package problems

import "container/heap"

type PairSum struct {
	s int // Primary sort key
	i int // Secondary sort key
}

var EmptyPair = PairSum{s: -1, i: -1}

// PSQueue implements heap.Interface and holds PairSums
type PSQueue []PairSum

// Len is part of sort.Interface
func (pq PSQueue) Len() int { return len(pq) }

// Less is part of sort.Interface.
// We want "smallest s, then smallest i", so we return true if i < j.
func (pq PSQueue) Less(i, j int) bool {
	if pq[i].s == pq[j].s {
		return pq[i].i < pq[j].i
	}
	return pq[i].s < pq[j].s
}

// Swap is part of sort.Interface
func (pq PSQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push is part of heap.Interface.
// Note: It uses a pointer receiver because it modifies the slice's length.
func (pq *PSQueue) Push(x interface{}) {
	item := x.(PairSum)
	*pq = append(*pq, item)
}

// Pop is part of heap.Interface.
// It removes the last element (which the heap package swapped to the end).
func (pq *PSQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Peek is a helper to see the top element without removing it.
func (pq PSQueue) Peek() PairSum {
	return pq[0]
}

func minimumPairRemoval3510(nums []int) int {
	n := len(nums)
	next := make([]int, n)
	prev := make([]int, n)
	sumAt := make([]PairSum, n)
	pq := &PSQueue{}
	heap.Init(pq)
	ans := 0
	bad := 0
	for i := range n {
		next[i] = i + 1
		prev[i] = i - 1
		if i+1 < n {
			if nums[i] > nums[i+1] {
				bad++
			}
			item := PairSum{s: nums[i] + nums[i+1], i: i}
			sumAt[i] = item
			heap.Push(pq, item)
		}
	}
	for bad > 0 {
		var top PairSum
		for pq.Len() > 0 {
			top = heap.Pop(pq).(PairSum)
			if sumAt[top.i] == top {
				break
			}
		}
		curId := top.i
		nextId := next[curId]
		prevId := prev[curId]
		next2Id := next[nextId]
		ps := nums[curId] + nums[nextId]
		sumAt[top.i] = EmptyPair
		if nums[curId] > nums[nextId] {
			bad--
		}
		if prevId >= 0 {
			if nums[prevId] > nums[curId] {
				bad--
			}
			if nums[prevId] > ps {
				bad++
			}
			sumAt[prevId] = EmptyPair
		}
		if next2Id < n {
			if nums[nextId] > nums[next2Id] {
				bad--
			}
			if ps > nums[next2Id] {
				bad++
			}
			sumAt[nextId] = EmptyPair
		}
		nums[curId] = ps
		next[curId] = next2Id
		if next2Id < n {
			prev[next2Id] = curId
		}
		if prevId >= 0 {
			np := PairSum{s: nums[prevId] + nums[curId], i: prevId}
			sumAt[prevId] = np
			heap.Push(pq, np)
		}
		if next2Id < n {
			np := PairSum{s: nums[curId] + nums[next2Id], i: curId}
			sumAt[curId] = np
			heap.Push(pq, np)
		}
		ans++
	}
	return ans
}
