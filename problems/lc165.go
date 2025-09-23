package problems

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	L := max(len(v1), len(v2))
	for i := range L {
		x1, x2 := 0, 0
		if i < len(v1) {
			x, _ := strconv.Atoi(v1[i])
			x1 = x
		}
		if i < len(v2) {
			x, _ := strconv.Atoi(v2[i])
			x2 = x
		}
		if x1 < x2 {
			return -1
		}
		if x1 > x2 {
			return 1
		}
	}
	return 0
}
