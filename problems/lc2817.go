package problems

import "github.com/emirpasic/gods/trees/redblacktree"

func minAbsoluteDifference(nums []int, x int) int {
	if x == 0 {
		return 0
	}
	N := len(nums)
	tree := redblacktree.NewWithIntComparator()

	res := abs(nums[0] - nums[N-1])

	for i := x; i < N; i++ {
		cur := nums[i]
		tree.Put(nums[i-x], nil)
		hi, hasHi := tree.Ceiling(cur)
		lo, hasLo := tree.Floor(cur)
		if hasHi {
			res = min(res, hi.Key.(int)-cur)
		}
		if hasLo {
			res = min(res, cur-lo.Key.(int))
		}
		if res == 0 {
			return 0
		}
	}
	return res
}
