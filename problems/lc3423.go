package problems

func maxAdjacentDistance(nums []int) int {
	N := len(nums)
	maxD := abs(nums[0] - nums[N-1])
	for i := 1; i < N; i++ {
		maxD = max(maxD, abs(nums[i]-nums[i-1]))
	}
	return maxD
}
