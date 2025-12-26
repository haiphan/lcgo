package problems

func bestClosingTime(customers string) int {
	n := len(customers)
	ry := 0
	for _, c := range customers {
		if c == 'Y' {
			ry++
		}
	}
	rn := n - ry
	p := n + 1
	ans := -1
	ln := 0
	for i, c := range customers {
		// ry + ln
		cp := ry + ln
		if cp < p {
			p = cp
			ans = i
		}
		if c == 'Y' {
			ry--
		} else {
			ln++
		}
	}
	if rn < p {
		ans = n
	}
	return ans
}
