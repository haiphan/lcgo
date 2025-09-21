package problems

const SB int = 2000001

func packetHash(p []int) int {
	return p[0] + SB*p[2]
}

type Dest struct {
	start int
	q     [][]int
	ps    map[int]bool
}

func DestContructor() *Dest {
	return &Dest{start: 0, q: make([][]int, 0), ps: make(map[int]bool)}
}

func (d *Dest) AddPacket(p []int) bool {
	k := packetHash(p)
	if d.ps[k] {
		return false
	}
	d.ps[k] = true
	d.q = append(d.q, p)
	return true
}

func (d *Dest) ForwardPacket() []int {
	if d.start == len(d.q) {
		return []int{}
	}
	top := d.q[d.start]
	k := packetHash(top)
	delete(d.ps, k)
	d.start++
	return top
}

func (d *Dest) CountLTE(target int) int {
	l, r := d.start, len(d.q)-1
	cnt := 0
	for l <= r {
		m := (l + r) >> 1
		if d.q[m][2] <= target {
			cnt = m - d.start + 1
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return cnt
}

func (d *Dest) GetCount(start, end int) int {
	if d.start == len(d.q) {
		return 0
	}
	return d.CountLTE(end) - d.CountLTE(start-1)
}

type Router struct {
	n, head, size int
	q             [][]int
	dToQ          map[int]*Dest
}

func RouterConstructor(memoryLimit int) Router {
	return Router{n: memoryLimit, head: 0, size: 0, q: make([][]int, memoryLimit), dToQ: make(map[int]*Dest)}
}

func (rt *Router) AddPacket(source int, destination int, timestamp int) bool {
	p := []int{source, destination, timestamp}
	dest, has := rt.dToQ[destination]
	if !has {
		dest = DestContructor()
		rt.dToQ[destination] = dest
	}
	ok := dest.AddPacket(p)
	if ok {
		top := rt.q[rt.head]
		i := (rt.head + rt.size) % rt.n
		rt.q[i] = p
		rt.size++
		if rt.size > rt.n {
			rt.size--
			rt.dToQ[top[1]].ForwardPacket()
			rt.head = (rt.head + 1) % rt.n
		}
	}
	return ok
}

func (rt *Router) ForwardPacket() []int {
	if rt.size == 0 {
		return []int{}
	}
	top := rt.q[rt.head]
	rt.head = (rt.head + 1) % rt.n
	rt.size--
	rt.dToQ[top[1]].ForwardPacket()
	return top
}

func (rt *Router) GetCount(destination int, startTime int, endTime int) int {
	dest, has := rt.dToQ[destination]
	if !has {
		return 0
	}
	return dest.GetCount(startTime, endTime)
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
