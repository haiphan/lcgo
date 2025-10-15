package problems

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	cur, prev := 1, 0
	ans := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			prev = cur
			cur = 1
		}
		ans = max(ans, cur>>1, min(cur, prev))
	}
	return ans
}
