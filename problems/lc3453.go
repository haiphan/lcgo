package problems

import "math"

func separateSquares(squares [][]int) float64 {
	// n := len(squares)
	const eps float64 = 1e-5
	y_min := squares[0][1]
	y_max := y_min
	total := 0
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		y_min = min(y_min, y)
		y_max = max(y_max, y+l)
		total += l * l
	}
	target_area := float64(total) / 2.0
	areaGte := func(y0 float64) bool {
		area_under := 0.0
		for _, sq := range squares {
			y, l := float64(sq[1]), float64(sq[2])
			if y0 > y {
				h := math.Min(y0-y, l)
				area_under += l * h
			}
		}
		return area_under >= target_area
	}
	low, high := float64(y_min), float64(y_max)
	for high-low >= eps {
		mid := (high + low) / 2.0
		if areaGte(mid) {
			high = mid
		} else {
			low = mid
		}
	}
	return high
}
