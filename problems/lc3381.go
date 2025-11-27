package problems

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	const BIG int = 1 << 48
	dp := make([]int, k)
	for i := 1; i < k; i++ {
		dp[i] = BIG
	}
	sum := 0
	ans := -BIG
	for i := 1; i <= n; i++ {
		sum += nums[i-1]
		ii := i % k
		if i >= k {
			ans = max(ans, sum-dp[ii])
		}
		dp[ii] = min(dp[ii], sum)
	}
	return int64(ans)
}
