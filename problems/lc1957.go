package problems

import "strings"

func makeFancyString(s string) string {
	N := len(s)
	var prev byte
	cnt := 0
	sb := strings.Builder{}
	for i := 0; i < N; i++ {
		if s[i] != prev {
			cnt = 1
			prev = s[i]
			sb.WriteByte(s[i])
			continue
		}
		if cnt < 2 {
			cnt++
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
