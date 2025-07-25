package problems

func maxSum(nums []int) int {
	N := len(nums)
	ns := make(map[int]bool, N)
	maxv := nums[0]
	sum := 0
	for _, x := range nums {
		if !ns[x] && x > 0 {
			sum += x
		}
		maxv = max(maxv, x)
		ns[x] = true
	}
	if maxv < 0 {
		return maxv
	}
	return sum
}
