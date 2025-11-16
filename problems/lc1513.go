package problems

func numSub(s string) int {
	n := len(s)
	cur, ans := 0, 0
	for i := 0; i <= n; i++ {
		if i == n || s[i] == '0' {
			ans = (ans + ((cur * (cur + 1)) >> 1)) % MOD
			cur = 0
		} else {
			cur++
		}
	}
	return ans
}
