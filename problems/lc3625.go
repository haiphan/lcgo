package problems

func compTwo(x int) int {
	if x < 2 {
		return 0
	}
	return x * (x - 1) / 2
}

func countTrapezoids3625(points [][]int) int {
	n := len(points)
	maxPairs := n * (n - 1) / 2
	slopes := make(map[[2]int]int, maxPairs)
	lines := make(map[[3]int]int, maxPairs)
	midPoints := make(map[[2]int]int, maxPairs)
	midLines := make(map[[5]int]int, maxPairs)
	for i := range n {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx, dy := x2-x1, y2-y1
			g := gcd(abs(dx), abs(dy))
			dx, dy = dx/g, dy/g
			if dx < 0 || (dx == 0 && dy < 0) {
				dx, dy = -dx, -dy
			}
			intercept := dx*y1 - dy*x1
			slopes[[2]int{dx, dy}]++
			lines[[3]int{dx, dy, intercept}]++
			midX, midY := x1+x2, y1+y2
			midPoints[[2]int{midX, midY}]++
			midLines[[5]int{dx, dy, intercept, midX, midY}]++
		}
	}
	ans := 0
	for _, c := range slopes {
		ans += compTwo(c)
	}
	for _, c := range lines {
		ans -= compTwo(c)
	}
	for _, c := range midPoints {
		ans -= compTwo(c)
	}
	for _, c := range midLines {
		ans += compTwo(c)
	}
	return ans
}
