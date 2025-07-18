package problems

import "math"

type lessThanOp func(a, b int) bool

func minCompare(a, b int) bool {
	return a < b
}

func maxCompare(a, b int) bool {
	return a > b
}

func heapPush(h []int, x int, lessFn lessThanOp) []int {
	h = append(h, x)
	i := len(h) - 1
	for i > 0 {
		p := getPa(i)
		if h[p] == h[i] || lessFn(h[p], h[i]) {
			break
		}
		h[p], h[i] = h[i], h[p]
		i = p
	}
	return h
}

func heapPop(heap []int, lessFn lessThanOp) []int {
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]
	i := 0
	l := i*2 + 1

	for l < len(heap) {
		j := getMinChild(heap, l, lessFn)

		if heap[j] == heap[i] || lessFn(heap[i], heap[j]) {
			return heap
		}
		heap[j], heap[i] = heap[i], heap[j]

		i = j
		l = i*2 + 1
	}

	return heap
}

func getMinChild(h []int, l int, lessFn lessThanOp) int {
	r := l + 1
	if r < len(h) && lessFn(h[r], h[l]) {
		return r
	}
	return l
}

func minimumDifference(nums []int) int64 {
	M := len(nums)
	N := M / 3
	N2 := N * 2
	maxH := make([]int, 0, M)
	minH := make([]int, 0, M)
	pre, post := make([]int, M), make([]int, M)
	cur := 0
	for i := range N {
		cur += nums[i]
		maxH = heapPush(maxH, nums[i], maxCompare)
		pre[i] = cur
	}
	for i := N; i < N2; i++ {
		top := maxH[0]
		if nums[i] < top {
			cur = cur - top + nums[i]
			maxH = heapPush(maxH, nums[i], maxCompare)
			maxH = heapPop(maxH, maxCompare)
		}
		pre[i] = cur
	}
	cur = 0
	for i := M - 1; i >= N2; i-- {
		cur += nums[i]
		minH = heapPush(minH, nums[i], minCompare)
		post[i] = cur
	}
	for i := N2 - 1; i >= N; i-- {
		top := minH[0]
		if top < nums[i] {
			cur += (nums[i] - top)
			minH = heapPush(minH, nums[i], minCompare)
			minH = heapPop(minH, minCompare)
		}
		post[i] = cur
	}
	res := math.MaxInt64
	for i := N - 1; i < N2; i++ {
		res = min(res, pre[i]-post[i+1])
	}
	return int64(res)
}
