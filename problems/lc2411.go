package problems

func smallestSubarrays(nums []int) []int {
	N := len(nums)
	var dp [30]int
	res := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		end := i
		for b := 0; b < 30; b++ {
			if (nums[i] & (1 << b)) != 0 {
				dp[b] = i
			}
			end = max(end, dp[b])
		}
		res[i] = end - i + 1
	}
	return res
}
