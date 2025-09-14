package problems

import "strings"

func rmVowels(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			c = '+'
		}
		sb.WriteRune(c)
	}
	return sb.String()
}

func spellchecker(wordlist []string, queries []string) []string {
	n := len(wordlist)
	ws := make(map[string]bool, n)
	cap := make(map[string]string, n)
	nov := make(map[string]string, n)
	for _, w := range wordlist {
		ws[w] = true
		low := strings.ToLower(w)
		rmv := rmVowels(low)
		if _, has := cap[low]; !has {
			cap[low] = w
		}
		if _, has := nov[rmv]; !has {
			nov[rmv] = w
		}
	}
	for i, q := range queries {
		if ws[q] {
			continue
		}
		low := strings.ToLower(q)
		if v, has := cap[low]; has {
			queries[i] = v
			continue
		}
		if v, has := nov[rmVowels(low)]; has {
			queries[i] = v
			continue
		}
		queries[i] = ""
	}
	return queries
}
