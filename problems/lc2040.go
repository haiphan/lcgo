package problems

import "slices"

func splitBySign(nums []int) ([]int, []int) {
	N := len(nums)
	pos := make([]int, 0, N)
	neg := make([]int, 0, N)
	for _, x := range nums {
		if x < 0 {
			neg = append(neg, -x)
		} else {
			pos = append(pos, x)
		}
	}
	slices.Reverse(neg)
	return neg, pos
}

func countLte(A []int, B []int, m int) int {
	cnt := 0
	bi := len(B) - 1
	for _, a := range A {
		for bi >= 0 && (B[bi]*a) > m {
			bi--
		}
		cnt += (bi + 1)
	}
	return cnt
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	A1, A2 := splitBySign(nums1)
	B1, B2 := splitBySign(nums2)
	negCnt := len(A1)*len(B2) + len(A2)*len(B1)
	sign := 1
	ki := int(k)
	if ki > negCnt {
		ki -= negCnt
	} else {
		ki = negCnt - ki + 1
		sign = -1
		B1, B2 = B2, B1
	}
	l, r := 0, int(1e10)
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
