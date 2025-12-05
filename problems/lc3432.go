package problems

func countPartitions3432(nums []int) int {
	sum := 0
	n := len(nums)
	for _, x := range nums {
		sum += x
	}
	left, ans := 0, 0
	for i := 0; i < n-1; i++ {
		left += nums[i]
		if (left*2-sum)%2 == 0 {
			ans++
		}
	}
	return ans
}
