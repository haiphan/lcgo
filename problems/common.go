package problems

import "lcgo/utils"

const MOD int = 1e9 + 7

type TreeNode = utils.TreeNode

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hPop(heap []int) []int {
	// swap top and bottom
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]

	// sort that top to bottom
	i := 0
	l, r := (i*2 + 1), (i*2 + 2)

	for l < len(heap) {
		j := minVal(l, r, heap)

		if heap[j] >= heap[i] {
			return heap
		}
		heap[j], heap[i] = heap[i], heap[j]

		i = j
		l, r = (i*2 + 1), (i*2 + 2)
	}

	return heap
}

func getPa(x int) int {
	return (x - 1) / 2
}

func getMinIndex(h []int, l int) int {
	r := l + 1
	if r < len(h) && h[r] < h[l] {
		return r
	}
	return l
}

func hPush(h []int, x int) []int {
	h = append(h, x)
	i := len(h) - 1
	for i > 0 {
		p := getPa(i)
		if h[p] <= h[i] {
			break
		}
		h[p], h[i] = h[i], h[p]
		i = p
	}
	return h
}

func intPow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := pow(x*x, n/2)
	if n%2 == 1 {
		res *= x
	}
	return res
}
