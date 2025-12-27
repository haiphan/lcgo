package problems

type SegmentTree struct {
	n    int
	tree []int
}

func NewSegmentTree(data []int) *SegmentTree {
	n := len(data)
	st := &SegmentTree{
		n:    n,
		tree: make([]int, 4*n),
	}
	st.build(data, 1, 0, n-1)
	return st
}

func (st *SegmentTree) build(data []int, node, start, end int) {
	if start == end {
		st.tree[node] = data[start]
		return
	}
	mid := (start + end) / 2
	st.build(data, 2*node, start, mid)
	st.build(data, 2*node+1, mid+1, end)
	st.tree[node] = min(st.tree[2*node], st.tree[2*node+1])
}

func (st *SegmentTree) Update(i, v int) {
	st.update(1, 0, st.n-1, i, v)
}

func (st *SegmentTree) update(node, start, end, idx, val int) {
	if start == end {
		st.tree[node] = val
		return
	}
	mid := (start + end) / 2
	if idx <= mid {
		st.update(2*node, start, mid, idx, val)
	} else {
		st.update(2*node+1, mid+1, end, idx, val)
	}
	st.tree[node] = min(st.tree[2*node], st.tree[2*node+1])
}

func (st *SegmentTree) GetLte(l, r, v int) int {
	return st.getLte(1, 0, st.n-1, l, r, v)
}

func (st *SegmentTree) getLte(node, start, end, l, r, v int) int {
	// 1. Range is completely outside OR the minimum in this range is > v
	if start > end || start > r || end < l || st.tree[node] > v {
		return -1
	}

	// 2. Leaf node found that satisfies the condition
	if start == end {
		return start
	}

	mid := (start + end) / 2

	// 3. Search Left Child First (to find the "first" index)
	res := st.getLte(2*node, start, mid, l, r, v)
	if res != -1 {
		return res
	}

	// 4. If not in left, search Right Child
	return st.getLte(2*node+1, mid+1, end, l, r, v)
}
func busiestServers(k int, arrival []int, load []int) []int {
	arr := make([]int, k)
	last := k - 1
	count := make([]int, k)
	st := NewSegmentTree(arr)
	maxc := 0
	for i, start := range arrival {
		ii := i % k
		sv := st.GetLte(ii, last, start)
		if sv == -1 && ii > 0 {
			sv = st.GetLte(0, ii-1, start)
		}
		if sv == -1 {
			continue
		}
		count[sv]++
		maxc = max(maxc, count[sv])
		st.Update(sv, start+load[i])
	}
	res := make([]int, 0, k)
	for i := range k {
		if count[i] == maxc {
			res = append(res, i)
		}
	}
	return res
}
