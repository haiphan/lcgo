package problems

func sumEven(x int) bool {
	s := 0
	for x > 0 {
		s += (x % 10)
		x /= 10
	}
	return s%2 == 0
}
func countEven(num int) int {
	if !sumEven(num) {
		num--
	}
	return num / 2
}
