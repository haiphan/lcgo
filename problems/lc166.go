package problems

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	n, d := numerator, denominator
	s := ""
	if n == 0 {
		return "0"
	}
	arr := make([]byte, 0)
	if (n < 0) != (d < 0) {
		s += "-"
	}
	n, d = abs(n), abs(d)
	ipart := n / d
	iStr := strconv.Itoa(ipart)
	s += iStr
	n %= d
	if n == 0 {
		return s
	}
	s += "."
	si := -1
	seen := make(map[int]int)
	for n > 0 {
		p, has := seen[n]
		if has {
			si = p
			break
		}
		seen[n] = len(arr)
		n *= 10
		arr = append(arr, byte((n/d)+'0'))
		n %= d
	}
	if si == -1 {
		return s + string(arr)
	}
	var sb strings.Builder
	for i, b := range arr {
		if si == i {
			sb.WriteByte('(')
		}
		sb.WriteByte(b)
	}
	sb.WriteByte(')')
	return s + sb.String()
}
