package problems

import "math"

func isPerfectSquare(x int) bool {
	sx := math.Sqrt(float64(x))
	return sx == math.Floor(sx)
}
func numSquares(n int) int {
	if n == 1 {
		return 1
	}
	sqn := math.Sqrt(float64(n))
	base := int(sqn)
	if base*base == n {
		return 1
	}
	for a := 1; a <= base; a++ {
		if isPerfectSquare(n - a*a) {
			return 2
		}
	}

	for a := 1; a <= base; a++ {
		da := a * a
		for b := a; b <= base; b++ {
			r := n - da - b*b
			if r < 0 {
				break
			}
			if isPerfectSquare(r) {
				return 3
			}
		}
	}
	return 4
}
