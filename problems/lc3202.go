package problems

func maximumLength2(nums []int, k int) int {
	N := len(nums)
	if k == 1 {
		return N
	}
	res := 2
	dp := make([]int, k)
	remains := make([]int, N)
	for i, v := range nums {
		remains[i] = v % k
	}
	for r := range k {
		for i := range k {
			dp[i] = 0
		}
		for _, cur := range remains {
			prev := (k + r - cur) % k
			dp[cur] = dp[prev] + 1
			res = max(res, dp[cur])
		}
	}
	return res
}
