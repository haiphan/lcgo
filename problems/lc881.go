package problems

func numRescueBoats(people []int, limit int) int {
	cnt := 0
	dp := make([]int, limit+1)
	for _, n := range people {
		dp[n]++
	}
	l, r := 0, limit
	for l <= r {
		for l < r && dp[l] == 0 {
			l++
		}
		for l < r && dp[r] == 0 {
			r--
		}
		if dp[l] <= 0 || dp[r] <= 0 {
			return cnt
		}
		if dp[l] > 0 && (l+r <= limit) {
			dp[l]--
		}
		dp[r]--
		cnt++
	}
	return cnt
}
