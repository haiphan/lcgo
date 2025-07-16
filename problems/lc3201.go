package problems

func maximumLength(nums []int) int {
	if len(nums) == 2 {
		return 2
	}
	even, odd, alt := 0, 0, 0
	prev := 2
	for _, x := range nums {
		cur := x & 1
		if cur == 1 {
			odd++
		} else {
			even++
		}
		if cur != prev {
			prev = cur
			alt++
		}
	}
	return max(odd, even, alt)
}
