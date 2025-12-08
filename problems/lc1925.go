package problems

import "math"

func IsPerfectSquare(n int) bool {
	if n == 0 || n == 1 {
		return true
	}

	rootFloat := math.Sqrt(float64(n))

	rootInt := int(rootFloat)

	return rootInt*rootInt == n
}
func countTriples(n int) int {
	c := 0
	ns := n * n
	for x := 1; x < n; x++ {
		xx := x * x
		for y := x + 1; y <= n; y++ {
			v := xx + y*y
			if v > ns {
				break
			}
			if IsPerfectSquare(v) {
				c += 2
			}
		}
	}
	return c
}
