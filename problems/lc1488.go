package problems

func avoidFlood(rains []int) []int {
	n := len(rains)
	lm := make(map[int]int)
	free := make([]int, 0)
	ans := make([]int, n)
	li := 0
	for i, r := range rains {
		ans[i] = 1
		if r == 0 {
			free = append(free, i)
		} else {
			ans[i] = -1
			p, has := lm[r]
			if has {
				// drain index
				di := -1
				for fi, fp := range free[li:] {
					if fp > p {
						di = li + fi
						break
					}
				}
				if di == -1 {
					return []int{}
				}
				ans[free[di]] = r
				free[di] = -1
				if di == li {
					for li < len(free) && free[li] == -1 {
						li++
					}
				}
			}
			lm[r] = i
		}
	}
	return ans
}
