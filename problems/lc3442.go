package problems

func maxDifference(s string) int {
	cm := [26]int{}
	N := len(s)
	for _, b := range s {
		cm[int(b-'a')]++
	}
	maxOdd, minEven := 0, N+1
	for _, x := range cm {
		if x == 0 {
			continue
		}
		if x%2 == 0 {
			minEven = min(minEven, x)
		} else {
			maxOdd = max(maxOdd, x)
		}
	}
	return maxOdd - minEven
}
