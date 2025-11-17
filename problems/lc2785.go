package problems

import "strings"

func isVowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func sortVowels(s string) string {
	cm := make([]int, 123)
	for i := range s {
		if isVowel(s[i]) {
			cm[s[i]]++
		}
	}
	var sb strings.Builder
	var vi byte = 0
	for i := range s {
		if isVowel(s[i]) {
			for vi < 123 && cm[vi] == 0 {
				vi++
			}
			cm[vi]--
			sb.WriteByte(vi)
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
