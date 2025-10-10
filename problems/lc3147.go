package problems

func maximumEnergy(energy []int, k int) int {
	ee := energy
	n := len(ee)
	dp := make([]int, k)
	maxv := -1001
	for i := n - 1; i >= n-k; i-- {
		dp[i%k] = ee[i]
		maxv = max(maxv, ee[i])
	}
	for i := n - k - 1; i >= 0; i-- {
		j := i % k
		dp[j] += ee[i]
		maxv = max(maxv, dp[j])
	}
	return maxv
}
