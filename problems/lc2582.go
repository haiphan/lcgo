package problems

func passThePillow(n int, time int) int {
	time = time % ((n - 1) * 2)
	if time+1 <= n {
		return time + 1
	}
	return n - (time % (n - 1))
}
