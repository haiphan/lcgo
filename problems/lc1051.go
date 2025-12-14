package problems

func heightChecker(heights []int) int {
	var cm [101]int
	for _, h := range heights {
		cm[h]++
	}
	cnt := 0
	i := 0
	for _, h := range heights {
		for cm[i] == 0 {
			i++
		}
		if i != h {
			cnt++
		}
		cm[i]--
	}
	return cnt
}
