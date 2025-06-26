package problems

func longestSubsequence(s string, k int) int {
	cur := 0
	N := len(s)
	steps := min(32, N)
	i, p := 0, 1
	for i < steps {
		j := N - 1 - i
		v := p * int(s[j]-'0')
		if cur+v > k {
			break
		} else {
			cur += v
		}
		i++
		p *= 2
	}
	c := 0
	for j := i; j < N; j++ {
		if s[N-1-j] == '1' {
			c++
		}
	}
	return N - c
}
