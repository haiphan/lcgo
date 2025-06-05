package problems

func maxSubarraySumCircular(nums []int) int {
	globalMax, globalMin := nums[0], nums[0]
	curMax, curMin := 0, 0
	total := 0
	for _, n := range nums {
		curMax = max(n, curMax+n)
		curMin = min(n, curMin+n)
		total += n
		globalMax = max(globalMax, curMax)
		globalMin = min(globalMin, curMin)
	}
	if globalMax > 0 {
		return max(globalMax, total-globalMin)
	}
	return globalMax
}
