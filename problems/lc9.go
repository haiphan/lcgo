package problems

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
