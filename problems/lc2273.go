package problems

func getStrHash(s string) int {
	var cm [26]int
	h := 17
	for _, c := range s {
		cm[c-'a']++
	}
	for _, v := range cm {
		h = (h*7 + v) & 0xffffffffff
	}
	return h
}

func removeAnagrams(words []string) []string {
	cur := -1
	stack := make([]string, 0, len(words))
	for _, w := range words {
		h := getStrHash(w)
		if h == cur {
			continue
		}
		stack = append(stack, w)
		cur = h
	}
	return stack
}
