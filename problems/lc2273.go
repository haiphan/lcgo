package problems

func getStrCount(s string) [26]int {
	var cm [26]int
	for _, c := range s {
		cm[c-'a']++
	}
	return cm
}

func removeAnagrams(words []string) []string {
	var cur [26]int
	stack := make([]string, 0, len(words))
	for _, w := range words {
		h := getStrCount(w)
		if h == cur {
			continue
		}
		stack = append(stack, w)
		cur = h
	}
	return stack
}
