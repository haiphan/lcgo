package problems

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	cur, prev := 0, 0
	ans := 1
	for i, v := range nums {
		cur++
		if i+1 == n || nums[i+1] <= v {
			ans = max(ans, cur>>1, min(cur, prev))
			prev = cur
			cur = 0
		}
	}
	return ans
}
