package problems

func possibleStringCount2(word string, k int) int {
	N, l, allCnt := len(word), 0, 1
	groups := make([]int, 0, N)
	for r := 0; r <= N; r++ {
		if r == N || word[r] != word[l] {
			c := r - l
			groups = append(groups, c)
			allCnt = allCnt * c % MOD
			l = r
		}
	}
	NG := len(groups)
	if k <= NG {
		return allCnt
	}
	dp := make([]int, k)
	ndp := make([]int, k)
	pre := make([]int, k)
	dp[0] = 1
	pre[0] = 1
	for i := range k {
		pre[i] = 1
	}
	for i := range NG {
		v := groups[i]
		ndp[0] = 0
		for j := 1; j < k; j++ {
			li, ri := j-v-1, j-1
			ls, rs := 0, pre[ri]
			if li >= 0 {
				ls = pre[li]
			}
			ndp[j] = (rs - ls + MOD) % MOD
		}
		dp, ndp = ndp, dp
		pre[0] = dp[0]
		for j := 1; j < k; j++ {
			pre[j] = (pre[j-1] + dp[j]) % MOD
		}
	}
	return (allCnt - pre[k-1] + MOD) % MOD
}
