package problems

func getStatus(a, b int) int {
	return ((a % 2) << 1) + (b % 2)
}

func calDiff(a, b, k int, s string) int {
	N := len(s)
	BIG := N + 1
	maxDiff := -BIG
	minVal := [4]int{BIG, BIG, BIG, BIG}
	ca, cb := 0, 0
	pa, pb := 0, 0
	l := -1
	for r := range N {
		cr := int(s[r] - '0')
		if cr == a {
			ca++
		} else if cr == b {
			cb++
		}
		for (r-l >= k) && (cb-pb >= 2) {
			lStatus := getStatus(pa, pb)
			minVal[lStatus] = min(minVal[lStatus], pa-pb)
			l++
			cl := int(s[l] - '0')
			if cl == a {
				pa++
			} else if cl == b {
				pb++
			}
		}
		rStatus := getStatus(ca, cb) ^ 2
		if minVal[rStatus] != BIG {
			maxDiff = max(maxDiff, ca-cb-minVal[rStatus])
		}
	}
	return maxDiff
}

func maxDifferenceII(s string, k int) int {
	N := len(s)
	BIG := N + 1
	maxDiff := -BIG
	for a := range 5 {
		for b := range 5 {
			if a == b {
				continue
			}
			maxDiff = max(maxDiff, calDiff(a, b, k, s))
		}
	}
	if maxDiff == -BIG {
		return 0
	}
	return maxDiff
}
