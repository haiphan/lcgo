package problems

import "strings"

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if strings.Index(op, "++") >= 0 {
			x++
		}
	}
	neg := len(operations) - x
	return x - neg
}
