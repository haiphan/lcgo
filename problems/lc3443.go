package problems

func maxDistance(s string, k int) int {
	N := len(s)
	if N == 1 {
		return 1
	}
	x, y := 0, 0
	maxD := 0
	for i, c := range s {
		switch c {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		default:
			x--
		}
		curMax := min(i+1, abs(x)+abs(y)+2*k)
		maxD = max(maxD, curMax)
	}
	return maxD
}
