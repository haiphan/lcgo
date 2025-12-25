package problems

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	n := len(happiness)
	ans := 0
	sort.Ints(happiness)
	d := 0
	lb := n - k
	for i := n - 1; i >= lb; i-- {
		v := happiness[i] - d
		if v <= 0 {
			break
		}
		ans += v
		d++
	}

	return int64(ans)
}
