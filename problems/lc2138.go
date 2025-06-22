package problems

import "strings"

func padEnd(s string, p byte, k int) string {
	need := k - len(s)
	if need <= 0 {
		return s
	}
	var sb strings.Builder
	sb.Grow(k)
	sb.WriteString(s)
	for i := 0; i < need; i++ {
		sb.WriteByte(p)
	}
	return sb.String()
}
func divideString(s string, k int, fill byte) []string {
	res := []string{}
	N := len(s)
	for i := 0; i < N; i += k {
		end := min(N, i+k)
		res = append(res, s[i:end])
	}
	res[len(res)-1] = padEnd(res[len(res)-1], fill, k)
	return res
}
