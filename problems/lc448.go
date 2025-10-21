package problems

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		v--
		if nums[v] > 0 {
			nums[v] = -nums[v]
		}
	}
	l := 0
	for i, v := range nums {
		if v > 0 {
			nums[l] = i + 1
			l++
		}
	}
	return nums[:l]
}
