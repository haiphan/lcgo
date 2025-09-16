package problems

func interchangeableRectangles(rectangles [][]int) int64 {
	rc := make(map[float64]int, len(rectangles))
	for _, rec := range rectangles {
		r := float64(rec[0]) / float64(rec[1])
		rc[r]++
	}
	res := 0
	for _, v := range rc {
		res += v * (v - 1) / 2
	}
	return int64(res)
}
