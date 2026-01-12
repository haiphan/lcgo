package problems

func minTimeToVisitAllPoints(points [][]int) int {
	t := 0
	x, y := points[0][0], points[0][1]
	for i := 1; i < len(points); i++ {
		nx, ny := points[i][0], points[i][1]
		dx, dy := abs(nx-x), abs(ny-y)
		dd := min(dx, dy)
		t += dx - dd + dy
		x, y = nx, ny
	}
	return t
}
