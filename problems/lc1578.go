package problems

func minCost1578(colors string, neededTime []int) int {
	ans := 0
	l := 0
	for r, rt := range neededTime {
		if r == l {
			continue
		}
		if colors[r] == colors[l] {
			if rt > neededTime[l] {
				ans += neededTime[l]
				l = r
			} else {
				ans += rt
			}
		} else {
			l = r
		}
	}
	return ans
}
