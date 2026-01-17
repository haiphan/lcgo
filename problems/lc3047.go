package problems

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	maxl := 0
	for i := range bottomLeft {
		x1, y1 := bottomLeft[i][0], bottomLeft[i][1]
		x2, y2 := topRight[i][0], topRight[i][1]
		for j := i + 1; j < n; j++ {
			x3, y3 := bottomLeft[j][0], bottomLeft[j][1]
			x4, y4 := topRight[j][0], topRight[j][1]
			xLeft := max(x1, x3)
			xRight := min(x2, x4)
			if xRight-xLeft <= maxl {
				continue
			}
			yBottom := max(y1, y3)
			yTop := min(y2, y4)
			if yTop-yBottom <= maxl {
				continue
			}
			maxl = min(xRight-xLeft, yTop-yBottom)
		}
	}
	return int64(maxl * maxl)
}
