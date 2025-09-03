package problems

import "sort"

const YMAX int = 1e9 + 1

func numberOfPairs2(points [][]int) int {
	n := len(points)
	res := 0
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] > points[j][0]
	})
	for i := range n {
		yi := points[i][1]
		ymax := YMAX
		for j := i + 1; j < n; j++ {
			yj := points[j][1]
			if yj >= yi && yj < ymax {
				ymax = yj
				res++
			}

		}
	}
	return res
}
