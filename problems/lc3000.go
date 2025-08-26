package problems

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxa, maxd := 0, 0
	for _, rec := range dimensions {
		a, b := rec[0], rec[1]
		d := a*a + b*b
		area := a * b
		if d > maxd {
			maxd = d
			maxa = area
		} else if d == maxd {
			maxa = max(maxa, area)
		}
	}

	return maxa
}
