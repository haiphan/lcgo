package problems

func maxDotProduct(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	const NEG = -1e6 - 1
	dp := make([]int, n2+1) // O(n₂) space

	// Base case: dp[0..n2] = NEG (for row 0)
	for j := range dp {
		dp[j] = NEG
	}

	for i := 0; i < n1; i++ {
		prev := dp[0] // dp[i-1][0] (always NEG)
		for j := 1; j <= n2; j++ {
			current := nums1[i] * nums2[j-1]
			old := dp[j] // Save dp[i-1][j] for next diagonal
			dp[j] = max(current, dp[j], dp[j-1], prev+current)
			prev = old // Update diagonal for next j
		}
	}
	return dp[n2]
}
