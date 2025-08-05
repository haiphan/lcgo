package problems

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	cnt := len(fruits)
	fmin := 1001
	for _, f := range fruits {
		if f >= fmin {
			continue
		}
		old := cnt
		for i, b := range baskets {
			if b >= f {
				cnt--
				baskets[i] = 0
				break
			}
		}
		if cnt == old {
			fmin = f
		}
	}
	return cnt
}
