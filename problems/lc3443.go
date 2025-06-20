package problems

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func maxDistance(s string, k int) int {
	N := len(s)
	if N == 1 {
		return 1
	}
	x, y := 0, 0
	maxD := 0
	for i, c := range s {
		d := i + 1
		if c == 'N' {
			y++
		} else if c == 'S' {
			y--
		} else if c == 'E' {
			x++
		} else {
			x--
		}
		curMax := min(d, intAbs(x)+intAbs(y)+2*k)
		maxD = max(maxD, curMax)
	}
	return maxD
}
