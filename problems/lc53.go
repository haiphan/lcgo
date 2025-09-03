package problems

func maxSubArray(nums []int) int {
	res := nums[0]
	prev := 0
	for _, x := range nums {
		prev += x
		res = max(res, prev)
		prev = max(prev, 0)
	}
	return res
}
