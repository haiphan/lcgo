package problems

import "slices"

var cntAns int = 0
var cntABC [3]int

type Fast2DTable struct {
	mask   uint32
	keysX  []int32
	keysY  []int32
	values []int32 // Using int32 since range is [-1, n]
}

func NewFast2DTable(n int) *Fast2DTable {
	// Size is power of 2, at least 2x the number of elements for speed
	size := uint32(1)
	for size < uint32(n*2) {
		size <<= 1
	}

	t := &Fast2DTable{
		mask:   size - 1,
		keysX:  make([]int32, size),
		keysY:  make([]int32, size),
		values: make([]int32, size),
	}

	// Initialize all values to -2 (our "Empty" sentinel)
	for i := range t.values {
		t.values[i] = -2
	}
	return t
}

// hash combines x and y into a single uint32 index
func (t *Fast2DTable) hash(x, y int32) uint32 {
	// A simpler, faster mixer for 32-bit coordinates
	h := uint32(x)*31 + uint32(y)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func (t *Fast2DTable) Put(x, y int32, val int32) {
	idx := t.hash(x, y) & t.mask
	for t.values[idx] != -2 {
		if t.keysX[idx] == x && t.keysY[idx] == y {
			t.values[idx] = val
			return
		}
		idx = (idx + 1) & t.mask
	}
	t.keysX[idx] = x
	t.keysY[idx] = y
	t.values[idx] = val
}

func (t *Fast2DTable) Get(x, y int32) (int32, bool) {
	idx := t.hash(x, y) & t.mask
	for t.values[idx] != -2 {
		if t.keysX[idx] == x && t.keysY[idx] == y {
			return t.values[idx], true
		}
		idx = (idx + 1) & t.mask
	}
	return -2, false
}

type SearchOp struct {
	numChar, ub int
	c1          byte
	c2          byte
}

func OneChar(s string) int {
	cnt := 1
	n := len(s)
	maxc := max(cntABC[0], cntABC[1], cntABC[2])
	for i := 1; i < n; i++ {
		if (cntAns == maxc) || (cnt+n-i <= cntAns) {
			break
		}
		if s[i] == s[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		cntAns = max(cntAns, cnt)
	}
	return cntAns
}
func TwoChars(s string, c1, c2 byte) int {
	ub := min(cntABC[c1-'a'], cntABC[c2-'a']) * 2
	if ub <= cntAns {
		return cntAns
	}
	n := len(s)
	sz := (n << 1) + 1
	offset := n
	m := 1
	pos := make([]int, sz)
	mark := make([]int, sz)
	pos[offset] = -1
	mark[offset] = m
	delta := 0
	for i := range s {
		if s[i] != c1 && s[i] != c2 {
			if n-(i+1) <= cntAns {
				return cntAns
			}
			m++
			pos[offset] = i
			mark[offset] = m
			delta = 0
			continue
		}
		if s[i] == c1 {
			delta++
		} else {
			delta--
		}
		idx := delta + offset
		if mark[idx] == m {
			cntAns = max(cntAns, i-pos[idx])
		} else {
			pos[idx] = i
			mark[idx] = m
		}
	}
	return cntAns
}

func ThreeChars(s string) int {
	ub := 3 * min(cntABC[0], cntABC[1], cntABC[2])
	if ub <= cntAns {
		return cntAns
	}
	cnt := [3]int{0, 0, 0}
	pos := NewFast2DTable(len(s))
	pos.Put(0, 0, -1)

	for i := range s {
		cnt[s[i]-'a']++

		x, y := int32(cnt[1]-cnt[0]), int32(cnt[2]-cnt[0])
		// key := [2]int{cnt[1] - cnt[0], cnt[2] - cnt[0]}

		if p, has := pos.Get(x, y); has {
			cntAns = max(cntAns, i-int(p))
		} else {
			pos.Put(x, y, int32(i))
		}
	}
	return cntAns
}

func longestBalanced(s string) int {
	cntAns = 1
	n := len(s)
	cntABC[0], cntABC[1], cntABC[2] = 0, 0, 0
	for _, c := range s {
		cntABC[c-'a']++
	}
	maxc := max(cntABC[0], cntABC[1], cntABC[2])
	if maxc == n || (cntABC[0] == cntABC[1] && cntABC[1] == cntABC[2]) {
		return n
	}
	ops := []SearchOp{
		{numChar: 1, ub: maxc},
		{numChar: 2, ub: min(cntABC[0], cntABC[1]) * 2, c1: 'a', c2: 'b'},
		{numChar: 2, ub: min(cntABC[0], cntABC[2]) * 2, c1: 'a', c2: 'c'},
		{numChar: 2, ub: min(cntABC[1], cntABC[2]) * 2, c1: 'b', c2: 'c'},
		{numChar: 3, ub: 3 * min(cntABC[0], cntABC[1], cntABC[2])},
	}
	slices.SortFunc(ops, func(a, b SearchOp) int {
		if a.ub == b.ub {
			return a.numChar - b.numChar
		}
		return b.ub - a.ub
	})
	for i := range ops {
		if ops[i].ub <= cntAns {
			break
		}
		if ops[i].numChar == 1 {
			OneChar(s)
		} else if ops[i].numChar == 2 {
			TwoChars(s, ops[i].c1, ops[i].c2)
		} else {
			ThreeChars(s)
		}
	}
	return cntAns
}
