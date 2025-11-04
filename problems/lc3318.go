package problems

import "math/bits"

const MAXNV = 51
const SHIFT = 6
const STEP = 64
const MASK = 63
const ZEROSTATE = (1 << MAXNV) - 1

type BitSetCnt struct {
	cc [MAXNV]uint64
	ns uint64
}

func NewBitSetCnt() *BitSetCnt {
	var cc [MAXNV]uint64
	cc[0] = ZEROSTATE
	return &BitSetCnt{cc: cc, ns: 1}
}

func (b *BitSetCnt) ShiftLeft(v, cur int) {
	i, bm := cur, uint64(1<<v)
	b.cc[i] &^= bm
	if b.cc[i] == 0 {
		b.ns &^= (1 << i)
	}
	i++
	b.cc[i] |= bm
	b.ns |= (1 << i)
}

func (b *BitSetCnt) ShiftRight(v, cur int) {
	i, bm := cur, uint64(1<<v)
	b.cc[i] &^= bm
	if b.cc[i] == 0 {
		b.ns &^= (1 << i)
	}
	i--
	b.cc[i] |= bm
	b.ns |= (1 << i)
}

func (b *BitSetCnt) XSum(x int) int {
	xsum, n, ns := 0, 0, b.ns
	for ns != 0 && n < x {
		freq := bits.Len64(ns) - 1
		vals := b.cc[freq]
		for vals != 0 && n < x {
			v := bits.Len64(vals) - 1
			xsum += freq * v
			vals &^= (1 << v)
			n++
		}
		if vals == 0 {
			ns &^= (1 << freq)
		}
	}
	return xsum
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	var freq [MAXNV]int
	bs := NewBitSetCnt()
	for i := range k {
		v := nums[i]
		bs.ShiftLeft(v, freq[v])
		freq[v]++
	}
	ans[0] = bs.XSum(x)
	for i := k; i < n; i++ {
		prev, cur := nums[i-k], nums[i]
		bs.ShiftRight(prev, freq[prev])
		freq[prev]--
		bs.ShiftLeft(cur, freq[cur])
		freq[cur]++
		ans[i-k+1] = bs.XSum(x)
	}
	return ans
}
