package problems

func smallestRepunitDivByK(k int) int {
	if k == 1 {
		return 1
	}
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	x, l := 1, 1
	for x > 0 {
		x = (x*10 + 1) % k
		l++
	}
	return l
}
