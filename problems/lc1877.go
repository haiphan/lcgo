package problems

func minPairSum(nums []int) int {
	// n := len(nums)
	var counts [100001]int
	l, r := 100001, 0
	for _, x := range nums {
		l = min(l, x)
		r = max(r, x)
		counts[x]++
	}
	ans := 0
	for l <= r {
		if counts[l] == 0 {
			l++
		} else if counts[r] == 0 {
			r--
		} else {
			ans = max(ans, l+r)
			counts[l]--
			counts[r]--
		}
	}
	return ans
}
