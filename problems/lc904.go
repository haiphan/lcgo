package problems

func totalFruit(fruits []int) int {
	cnt, lastCnt, res := 0, 0, 0
	// p2, p1, cur
	p1, p2 := -1, -1
	for _, f := range fruits {
		if f == p1 || f == p2 {
			cnt++
		} else {
			cnt = lastCnt + 1
		}
		if f == p1 {
			lastCnt++
		} else {
			p2, p1, lastCnt = p1, f, 1
		}
		res = max(res, cnt)
	}
	return res
}
