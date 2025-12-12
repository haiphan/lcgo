package problems

import "strings"

func commonChars(words []string) []string {
	n := len(words)
	if n == 1 {
		return strings.Split(words[0], "")
	}
	var cm [26]int
	w := words[0]
	m := len(w)
	for i := range words[0] {
		cm[w[i]-'a']++
	}
	for i := 1; i < n; i++ {
		w := words[i]
		var cc [26]int
		for j := range w {
			cc[w[j]-'a']++
		}
		for j, v := range cc {
			if v < cm[j] {
				cm[j] = v
			}
		}
	}
	ans := make([]string, 0, m)
	for i, v := range cm {
		for v > 0 {
			ans = append(ans, string(rune(i+'a')))
			v--
		}
	}
	return ans
}
