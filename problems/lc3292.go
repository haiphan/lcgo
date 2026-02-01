package problems

func prefix_function(s, t string) []int {
	n := len(s) + 1 + len(t) // s + "#" + t
	pi := make([]int, n)
	
	getChar := func(i int) byte {
		if i < len(s) {
			return s[i]
		} else if i == len(s) {
			return '#'
		} else {
			return t[i-len(s)-1]
		}
	}
	
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && getChar(i) != getChar(j) {
			j = pi[j-1]
		}
		if getChar(i) == getChar(j) {
			j++
		}
		pi[i] = j
	}
	return pi
}
func minValidStrings(words []string, target string) int {
	n := len(target)
	ps := make([]int, n+1)
	for _, w := range words {
		pi := prefix_function(w, target)
		for i := 1; i <= n; i++ {
			ps[i] = max(ps[i], pi[len(w)+i])
		}
	}
	len, ans := n, 0
	for ; len > 0 && ps[len] > 0; ans++ {
		len -= ps[len]
	}
	if len == 0 {
		return ans
	}
	return -1
}
