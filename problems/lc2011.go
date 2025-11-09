package problems

import "strings"

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if strings.Contains(op, "++") {
			x++
		}
	}
	neg := len(operations) - x
	return x - neg
}
