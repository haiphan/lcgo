package problems

func minBitwiseArray3315(nums []int) []int {
	for i, x := range nums {
		if x&1 == 0 {
			nums[i] = -1
		} else {
			nums[i] = x &^ (((x + 1) &^ x) >> 1)
		}
	}

	return nums
}
