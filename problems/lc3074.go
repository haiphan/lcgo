package problems

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for _, a := range apple {
		sum += a
	}
	sort.Ints(capacity)
	n := len(capacity)
	for i := n - 1; i >= 0; i-- {
		sum -= capacity[i]
		if sum <= 0 {
			return n - i
		}
	}
	return n
}
