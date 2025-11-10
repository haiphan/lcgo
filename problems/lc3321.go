package problems

import "container/heap"

type FItem struct {
	f, v, i int
	x       bool
}

func compareBig(a, b *FItem) bool {
	if a.f == b.f {
		return a.v > b.v
	}
	return a.f > b.f
}

func compareSmall(a, b *FItem) bool {
	return !compareBig(a, b)
}

type CompareFn func(*FItem, *FItem) bool

type FreqHeap struct {
	q      []*FItem
	lessFn CompareFn
}

func NewFHeap(f CompareFn) *FreqHeap {
	return &FreqHeap{q: []*FItem{}, lessFn: f}
}

func (h FreqHeap) Len() int { return len(h.q) }

func (h FreqHeap) Less(i, j int) bool {
	return h.lessFn(h.q[i], h.q[j])
}

func (h FreqHeap) Swap(i, j int) {
	h.q[i], h.q[j] = h.q[j], h.q[i]
	h.q[i].i = i
	h.q[j].i = j
}

func (h *FreqHeap) Push(x any) {
	n := len(h.q)
	item := x.(*FItem)
	h.q = append(h.q, item)
	item.i = n
}

func (h *FreqHeap) Pop() any {
	old := h.q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.i = -1
	item.x = false
	h.q = old[:n-1]
	return item
}

func (h *FreqHeap) RemoveItem(x *FItem) *FItem {
	if x.i < 0 || x.i >= len(h.q) || h.q[x.i] != x {
		return nil
	}
	removedItem := heap.Remove(h, x.i)
	return removedItem.(*FItem)
}

func (h *FreqHeap) UpdateItem(x *FItem, freq int) {
	if x.i < 0 || x.i >= len(h.q) || h.q[x.i] != x {
		return
	}

	x.f = freq
	if freq == 0 {
		h.RemoveItem(x)
		return
	}
	heap.Fix(h, x.i)
}

func (h *FreqHeap) Peek() *FItem {
	if h.Len() == 0 {
		return nil
	}
	return h.q[0]
}

func findXSum3321(nums []int, k int, x int) []int64 {
	xHeap := NewFHeap(compareSmall)
	rHeap := NewFHeap(compareBig)
	n := len(nums)
	nc := make(map[int]*FItem)
	sum := 0
	ans := make([]int64, n-k+1)
	for i, cur := range nums {
		curItem, has := nc[cur]
		if !has {
			curItem = &FItem{f: 1, v: cur, i: -1, x: false}
			nc[cur] = curItem
			heap.Push(rHeap, curItem)
		} else if curItem.x {
			sum += cur
			xHeap.UpdateItem(curItem, curItem.f+1)
		} else {
			oldF := curItem.f
			curItem.f++
			if oldF == 0 {
				heap.Push(rHeap, curItem)
			} else {
				rHeap.UpdateItem(curItem, oldF+1)
			}
		}
		if i >= k {
			prev := nums[i-k]
			item := nc[prev]
			f := item.f - 1
			if item.x {
				xHeap.UpdateItem(item, f)
				sum -= prev
			} else {
				rHeap.UpdateItem(item, f)
			}
		}

		if rHeap.Len() > 0 && xHeap.Len() < x {
			item := heap.Pop(rHeap).(*FItem)
			item.x = true
			heap.Push(xHeap, item)
			sum += item.f * item.v
		}
		if rHeap.Len() > 0 && xHeap.Len() == x && compareSmall(xHeap.Peek(), rHeap.Peek()) {
			rItem := heap.Pop(rHeap).(*FItem)
			xItem := heap.Pop(xHeap).(*FItem)
			rItem.x = true
			sum += rItem.f*rItem.v - xItem.f*xItem.v
			heap.Push(xHeap, rItem)
			heap.Push(rHeap, xItem)
		}
		if i >= k-1 {
			ans[i-k+1] = int64(sum)
		}
	}
	return ans
}
