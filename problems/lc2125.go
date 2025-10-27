package problems

import "strings"

func numberOfBeams(bank []string) int {
	prev := 0
	ans := 0
	for _, row := range bank {
		cur := strings.Count(row, "1")
		if cur == 0 {
			continue
		}
		ans += cur * prev
		prev = cur
	}
	return ans
}
