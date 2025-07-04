package problems

import (
	"sort"
)

const MAXP int = 1e10

func revAndAbs(A []int) []int {
	N := len(A)
	for i := range N {
		A[i] = -A[i]
	}
	for l, r := 0, N-1; l < r; l, r = l+1, r-1 {
		A[l], A[r] = A[r], A[l]
	}
	return A
}

func countLte(A []int, B []int, m int) int {
	cnt := 0
	bi := len(B) - 1
	for _, a := range A {
		for bi >= 0 && (B[bi]*a) > m {
			bi--
		}
		if bi < 0 {
			return cnt
		}
		cnt += (bi + 1)
	}
	return cnt
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	zA := sort.SearchInts(nums1, 0)
	zB := sort.SearchInts(nums2, 0)
	A1, A2 := revAndAbs(nums1[:zA]), nums1[zA:]
	B1, B2 := revAndAbs(nums2[:zB]), nums2[zB:]
	negCnt := len(A1)*len(B2) + len(A2)*len(B1)
	sign := 1
	k32 := int(k)
	ki := k32 - negCnt
	if k32 <= negCnt {
		ki = negCnt - k32 + 1
		sign = -1
		B1, B2 = B2, B1
	}
	l, r := 0, MAXP
	for l < r {
		m := (l + r) / 2
		if countLte(A1, B1, m)+countLte(A2, B2, m) >= ki {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(sign * l)
}
