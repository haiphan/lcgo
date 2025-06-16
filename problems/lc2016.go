package problems

func maximumDifference(nums []int) int {
	minV, N, res := nums[0], len(nums), -1
	for i := 1; i < N; i++ {
		if nums[i] > minV {
			res = max(res, nums[i]-minV)
		} else {
			minV = nums[i]
		}
	}
	return res
}
