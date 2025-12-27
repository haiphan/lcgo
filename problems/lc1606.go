package problems

import "container/heap"

type Server struct {
	a, b int
}

type ServerHeap []Server

func (sh ServerHeap) Len() int { return len(sh) }

func (sh ServerHeap) Less(i, j int) bool {
	return sh[i].a < sh[j].a
}

func (sh ServerHeap) Swap(i, j int) {
	sh[i], sh[j] = sh[j], sh[i]
}

func (sh *ServerHeap) Push(x any) {
	*sh = append(*sh, x.(Server))
}

func (sh *ServerHeap) Pop() any {
	old := *sh
	n := len(old)
	x := old[n-1]
	*sh = old[0 : n-1]
	return x
}

func busiestServers(k int, arrival []int, load []int) []int {
	count := make([]int, k)
	free := make([]Server, k)
	busy := make([]Server, 0)
	for i := range k {
		free[i] = Server{a: i, b: 0}
	}
	freeHeap := (*ServerHeap)(&free)
	busyHeap := (*ServerHeap)(&busy)
	heap.Init(freeHeap)
	heap.Init(busyHeap)
	maxc := 0
	for i, t := range arrival {
		for busyHeap.Len() > 0 && (*busyHeap)[0].a <= t {
			top := heap.Pop(busyHeap).(Server)
			s := top.b
			if s < i {
				s = i + ((s-i)%k+k)%k
			}
			top.a = s
			heap.Push(freeHeap, top)
		}
		if freeHeap.Len() == 0 {
			continue
		}
		top := heap.Pop(freeHeap).(Server)
		si := top.a % k
		count[si]++
		maxc = max(maxc, count[si])
		top.a, top.b = t+load[i], top.a
		heap.Push(busyHeap, top)
	}
	res := make([]int, 0)
	for i := range k {
		if count[i] == maxc {
			res = append(res, i)
		}
	}
	return res
}
