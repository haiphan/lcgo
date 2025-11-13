package problems

func maxOperations(s string) int {
	c1, ans := 0, 0
	for i := range s {
		c1 += int(s[i] - '0')
		if i > 0 && s[i] < s[i-1] {
			ans += c1
		}
	}
	return ans
}
