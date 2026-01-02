package problems

func longestMonotonicSubarray(nums []int) int {
	dec, inc := 1, 1
	ci, cd := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ci, cd = 1, 1
		} else if nums[i] > nums[i-1] {
			ci++
			cd = 1
		} else {
			cd++
			ci = 1
		}
		dec = max(dec, cd)
		inc = max(inc, ci)
	}
	return max(dec, inc)
}
