package problems

import "math"

func getTriangleArea(a, b, c []int) float64 {
	return math.Abs(float64((b[0]-a[0])*(c[1]-a[1])-(b[1]-a[1])*(c[0]-a[0]))) / 2.0
}
func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	res := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				res = math.Max(res, getTriangleArea(points[i], points[j], points[k]))
			}
		}
	}
	return res
}
