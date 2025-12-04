package problems

func getLastMoment(n int, left []int, right []int) int {
	t := 0
	for _, x := range left {
		t = max(x, t)
	}
	for _, x := range right {
		t = max(t, n-x)
	}
	return t
}
