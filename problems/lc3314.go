package problems

func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if (x & 1) == 1 {
			nums[i] = (x &^ (((x + 1) &^ x) >> 1))
		} else {
			nums[i] = -1
		}
	}
	return nums
}
