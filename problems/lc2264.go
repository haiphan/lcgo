package problems

import "strings"

var cand []string = []string{"999", "888", "777", "666", "555", "444", "333", "222", "111", "000"}

func largestGoodInteger(num string) string {
	for _, c := range cand {
		if strings.Contains(num, c) {
			return c
		}
	}
	return ""
}
