package problems

func sumZero(n int) []int {
	res := make([]int, 0)
	if (n & 1) == 1 {
		n--
		res = append(res, 0)
	}
	for ; n > 0; n -= 2 {
		res = append(res, n)
		res = append(res, -n)
	}
	return res
}
