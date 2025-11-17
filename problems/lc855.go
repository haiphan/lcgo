package problems

import "container/heap"

type SeatRange struct {
	Head, Tail, Index int
}

type SeatCompareFn func(*SeatRange, *SeatRange) bool

type SRHeap struct {
	q      []*SeatRange
	lessFn SeatCompareFn
}

func NewSRHeap(f SeatCompareFn) *SRHeap {
	return &SRHeap{q: []*SeatRange{}, lessFn: f}
}

func (h SRHeap) Len() int { return len(h.q) }

func (h SRHeap) Less(i, j int) bool {
	return h.lessFn(h.q[i], h.q[j])
}

func (h SRHeap) Swap(i, j int) {
	h.q[i], h.q[j] = h.q[j], h.q[i]
	h.q[i].Index = i
	h.q[j].Index = j
}

func (h *SRHeap) Push(x any) {
	n := len(h.q)
	item := x.(*SeatRange)
	h.q = append(h.q, item)
	item.Index = n
}

func (h *SRHeap) Pop() any {
	old := h.q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	h.q = old[:n-1]
	return item
}

func (h *SRHeap) RemoveItem(x *SeatRange) *SeatRange {
	if x.Index < 0 || x.Index >= len(h.q) || h.q[x.Index] != x {
		return nil
	}
	removedItem := heap.Remove(h, x.Index)
	return removedItem.(*SeatRange)
}

func (h *SRHeap) Peek() *SeatRange {
	if h.Len() == 0 {
		return nil
	}
	return h.q[0]
}

type ExamRoom struct {
	q       *SRHeap
	n       int
	headmap map[int]*SeatRange
	tailmap map[int]*SeatRange
}

func getLen(l, r, n int) int {
	if l == 0 {
		return r
	}
	if r == n-1 {
		return r - l
	}
	return (r - l) >> 1
}

func ExamRoomConstructor(n int) ExamRoom {
	lessFn := func(a, b *SeatRange) bool {
		d1 := getLen(a.Head, a.Tail, n)
		d2 := getLen(b.Head, b.Tail, n)
		if d1 == d2 {
			return a.Head < b.Head
		}
		return d1 > d2
	}
	q := NewSRHeap(lessFn)
	eRoom := ExamRoom{q: q, n: n, headmap: make(map[int]*SeatRange), tailmap: make(map[int]*SeatRange)}
	eRoom.PushRange(0, n-1)
	return eRoom
}

func (eRoom *ExamRoom) PushRange(l, r int) {
	sr := &SeatRange{Head: l, Tail: r}
	heap.Push(eRoom.q, sr)
	eRoom.headmap[l] = sr
	eRoom.tailmap[r] = sr
}

func (eRoom *ExamRoom) RemoveRange(sr *SeatRange) {
	if sr.Index >= 0 {
		heap.Remove(eRoom.q, sr.Index)
	}
	delete(eRoom.headmap, sr.Head)
	delete(eRoom.tailmap, sr.Tail)
}

func (eRoom *ExamRoom) Seat() int {
	sr := heap.Pop(eRoom.q).(*SeatRange)
	eRoom.RemoveRange(sr)
	var p int
	if sr.Head == 0 {
		p = 0
	} else if sr.Tail == eRoom.n-1 {
		p = eRoom.n - 1
	} else {
		p = (sr.Head + sr.Tail) >> 1
	}
	if p > sr.Head {
		eRoom.PushRange(sr.Head, p-1)
	}
	if p < sr.Tail {
		eRoom.PushRange(p+1, sr.Tail)
	}
	return p
}

func (eRoom *ExamRoom) Leave(p int) {
	l, r := p, p
	lRange := eRoom.tailmap[p-1]
	rRange := eRoom.headmap[p+1]
	if lRange != nil {
		l = lRange.Head
		eRoom.RemoveRange(lRange)
	}
	if rRange != nil {
		r = rRange.Tail
		eRoom.RemoveRange(rRange)
	}
	eRoom.PushRange(l, r)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
