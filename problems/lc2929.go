package problems

func getComb(n int) int {
	if n < 0 {
		return 0
	}
	return (n + 2) * (n + 1) / 2
}
func distributeCandies(n int, limit int) int64 {
	l := limit
	l1 := l + 1
	return int64(getComb(n) - 3*getComb(n-l1) + 3*getComb(n-2*l1) - getComb(n-3*l1))
}
