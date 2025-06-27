package problems

func countSub(s string, t string) int {
	SL, TL := len(s), len(t)
	i, j, cnt := 0, 0, 0
	for i < SL {
		if s[i] == t[j] {
			j++
			if j == TL {
				cnt++
				j = 0
			}
		}
		i++
	}
	return cnt
}

func longestSubsequenceRepeatedK(s string, k int) string {
	cc := [26]int{}
	for i := range len(s) {
		cc[s[i]-'a']++
	}
	qi := 0
	q := []string{""}
	res := ""
	for qi < len(q) {
		cur := q[qi]
		qi++
		for i := range 26 {
			if cc[i] < k {
				continue
			}
			next := cur + string(byte(i)+'a')
			if countSub(s, next) >= k {
				res = next
				q = append(q, next)
			}
		}
	}
	return res
}
