package problems

func findAnagrams(s string, p string) []int {
	L := len(p)
	ans := []int{}
	var cm [26]int
	for _, c := range p {
		cm[c-'a']++
	}
	var cc [26]int
	l := 0
	for r, c := range s {
		code := int(c - 'a')
		cc[code]++
		for l <= r && cc[code] > cm[code] {
			cc[s[l]-'a']--
			l++
		}
		if r-l+1 == L {
			ans = append(ans, l)
		}
	}
	return ans
}
