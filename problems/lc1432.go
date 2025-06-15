package problems

import (
	"strconv"
	"strings"
)

func getReplace(num int, isMax bool) int {
	s := strconv.Itoa(num)
	old := byte('a')
	for i := range len(s) {
		c := s[i]
		if isMax {
			if c != '9' {
				old = c
				break
			}
		} else {
			if c != '0' && c != '1' {
				old = c
				break
			}
		}
	}
	if old == 'a' {
		return num
	}
	nc := '9'
	if !isMax {
		nc = '0'
		if s[0] == old {
			nc = '1'
		}
	}
	s = strings.ReplaceAll(s, string(old), string(nc))
	x, _ := strconv.Atoi(s)
	return x
}
func maxDiff(num int) int {
	return getReplace(num, true) - getReplace(num, false)
}
