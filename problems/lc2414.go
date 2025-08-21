package problems

func longestContinuousSubstring(s string) int {
	N := len(s)
	res, cnt := 1, 1
	for i := 1; i < N; i++ {
		if s[i]-s[i-1] == 1 {
			cnt++
		} else {
			cnt = 1
		}
		res = max(res, cnt)
	}
	return res
}
