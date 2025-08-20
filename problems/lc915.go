package problems

func partitionDisjoint(nums []int) int {
	n := len(nums)
	lmax, gmax := nums[0], nums[0]
	res := 0
	for i := 1; i < n; i++ {
		if gmax < nums[i] {
			gmax = nums[i]
		}
		if lmax > nums[i] {
			lmax = gmax
			res = i
		}
	}
	return res + 1
}
