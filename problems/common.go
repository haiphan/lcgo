package problems

import (
	"lcgo/utils"
	"math"
)

const MOD int = 1e9 + 7

type TreeNode = utils.TreeNode

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func rvInt(x int) int {
	if x == 0 {
		return 0
	}
	y := 0
	for x > 0 {
		y = y*10 + (x % 10)
		x /= 10
	}
	return y
}

func getHead(x, l int) int {
	if x == 0 {
		return 0
	}

	numDigits := int(math.Log10(float64(x))) + 1

	divisor := int(math.Pow10(numDigits - l))

	return x / divisor
}

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
		c := l
		if r < len(heap) && heap[r] < heap[l] {
			c = r
		}

		if heap[c] >= heap[i] {
			return heap
		}
		heap[c], heap[i] = heap[i], heap[c]

		i = c
		l, r = (i*2 + 1), (i*2 + 2)
	}

	return heap
}

func hPush(h []int, x int) []int {
	h = append(h, x)
	i := len(h) - 1
	for i > 0 {
		p := (i - 1) >> 1
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

func intPowMod(x, n, m int) int {
	if n == 0 {
		return 1 % m
	}
	if n == 1 {
		return x % m
	}
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % m
		}
		x = (x * x) % m
		n = n >> 1
	}
	return res
}
