package problems

func findClosest(x int, y int, z int) int {
	d1, d2 := abs(x-z), abs(y-z)
	if d1 == d2 {
		return 0
	}
	if d1 < d2 {
		return 1
	}
	return 2
}
