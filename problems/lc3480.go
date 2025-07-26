package problems

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	right := make([][]int, n+1)
	bonus := make([]int, n+1)
	res, maxBonus := 0, 0
	l0, l1 := 0, 0
	// l0, l1, r. l0 is the biggest
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		right[b] = append(right[b], a)
	}

	for r := 1; r <= n; r++ {
		for _, l := range right[r] {
			if l > l0 {
				l1 = l0
				l0 = l
			} else if l > l1 {
				l1 = l
			}
		}
		res += r - l0
		if l0 > 0 {
			bonus[l0] += l0 - l1
			maxBonus = max(maxBonus, bonus[l0])
		}
	}
	return int64(res + maxBonus)
}
