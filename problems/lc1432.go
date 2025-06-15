package problems

import (
	"strconv"
	"strings"
)

func checkMin(c byte) bool {
	return c != '0' && c != '1'
}

func checkMax(c byte) bool {
	return c != '9'
}
func getReplace(num int, isMax bool) int {
	s := strconv.Itoa(num)
	old := byte('a')
	chkFn := checkMin
	if isMax {
		chkFn = checkMax
	}
	for i := range len(s) {
		c := s[i]
		if chkFn(c) {
			old = c
			break
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
