package problems

import "math/bits"

func countPalindromicSubsequence(s string) int {
	ans := 0
	var first [26]int
	var last [26]int
	for i := range first {
		first[i] = -1
	}
	for i := range s {
		code := int(s[i] - 'a')
		if first[code] == -1 {
			first[code] = i
		}
		last[code] = i
	}
	for i := range 26 {
		head, tail := first[i], last[i]
		if head >= 0 && head < tail {
			var cur uint32
			for j := head + 1; j < tail; j++ {
				cur |= (1 << uint32(s[j]-'a'))
			}
			ans += bits.OnesCount32(cur)
		}
	}
	return ans
}
