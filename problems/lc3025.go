package problems

import "sort"

const BIGV int = 51

func numberOfPairs(points [][]int) int {
	n := len(points)
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] > points[j][0]
	})
	res := 0
	for i := range n {
		yi, ymax := points[i][1], BIGV
		for j := i + 1; j < n; j++ {
			yj := points[j][1]
			if yj >= yi && yj < ymax {
				res++
				ymax = yj
			}
		}
	}
	return res
}
