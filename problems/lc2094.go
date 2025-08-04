package problems

func isGood(n int, dc []int) bool {
	d1, d10, d100 := n%10, (n/10)%10, n/100
	dc[d1]--
	dc[d10]--
	dc[d100]--
	good := dc[d1] >= 0 && dc[d10] >= 0 && dc[d100] >= 0
	dc[d1]++
	dc[d10]++
	dc[d100]++
	return good
}

func findEvenNumbers(digits []int) []int {
	dc := make([]int, 10)
	for _, d := range digits {
		dc[d]++
	}
	res := make([]int, 0)
	for i := 100; i < 1000; i += 2 {
		if isGood(i, dc) {
			res = append(res, i)
		}
	}
	return res
}
