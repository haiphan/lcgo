package problems

import "math/bits"

type Dest struct {
	start int
	q     []int
}

func DestContructor() *Dest {
	return &Dest{start: 0, q: make([]int, 0)}
}

func (d *Dest) AddPacket(p int) {
	d.q = append(d.q, p)
}

func (d *Dest) ForwardPacket() {
	if d.start == len(d.q) {
		return
	}
	d.start++
	if d.start == len(d.q) {
		d.q = make([]int, 0)
		d.start = 0
	}
}

func (d *Dest) CountLTE(target int) int {
	l, r := d.start, len(d.q)
	for l < r {
		m := (l + r) >> 1
		if d.q[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - d.start
}

func (d *Dest) GetCount(start, end int) int {
	if d.start == len(d.q) {
		return 0
	}
	endCnt := d.CountLTE(end)
	if endCnt == 0 {
		return 0
	}
	return endCnt - d.CountLTE(start-1)
}

type Router struct {
	n, head, size int
	q             [][3]int
	pm            map[[3]int]bool
	dToQ          map[int]*Dest
	mask          int
}

func RouterConstructor(memoryLimit int) Router {
	cap := memoryLimit
	if memoryLimit&(memoryLimit-1) != 0 {
		cap = 1 << int(bits.Len(uint(memoryLimit-1)))
	}
	return Router{
		n:    memoryLimit,
		pm:   make(map[[3]int]bool),
		head: 0, size: 0,
		q:    make([][3]int, cap),
		dToQ: make(map[int]*Dest),
		mask: cap - 1,
	}
}

func (rt *Router) AddPacket(source int, destination int, timestamp int) bool {
	pk := [3]int{source, destination, timestamp}
	if rt.pm[pk] {
		return false
	}
	rt.pm[pk] = true
	if rt.size == rt.n {
		rt.ForwardPacket()
	}
	i := (rt.head + rt.size) & rt.mask
	rt.q[i] = pk
	rt.size++
	dest, has := rt.dToQ[destination]
	if !has {
		dest = DestContructor()
		rt.dToQ[destination] = dest
	}
	dest.AddPacket(timestamp)
	return true
}

func (rt *Router) ForwardPacket() []int {
	if rt.size == 0 {
		return []int{}
	}
	top := rt.q[rt.head]
	delete(rt.pm, top)
	rt.head = (rt.head + 1) & rt.mask
	rt.size--
	rt.dToQ[top[1]].ForwardPacket()
	return top[:]
}

func (rt *Router) GetCount(destination int, startTime int, endTime int) int {
	if startTime > endTime {
		return 0
	}
	dest, has := rt.dToQ[destination]
	if has {
		return dest.GetCount(startTime, endTime)
	}
	return 0
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
