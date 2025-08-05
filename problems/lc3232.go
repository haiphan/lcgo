package problems

func canAliceWin(nums []int) bool {
	t := 0
	for _, x := range nums {
		if x < 10 {
			t += x
		} else if x < 100 {
			t -= x
		}
	}
	return t != 0
}
