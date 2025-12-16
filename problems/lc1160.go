package problems

func countCharacters(words []string, chars string) int {
	var cm [26]int
	n := len(chars)
	for _, c := range chars {
		cm[c-'a']++
	}
	ans := 0
	for _, w := range words {
		if len(w) > n {
			continue
		}
		var cc [26]int
		skip := false
		for _, c := range w {
			cc[c-'a']++
			if cc[c-'a'] > cm[c-'a'] {
				skip = true
				break
			}
		}
		if !skip {
			ans += len(w)
		}
	}
	return ans
}
