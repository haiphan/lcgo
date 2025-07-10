package problems

func deleteAndEarn(nums []int) int {
	maxv := 0
	for _, v := range nums {
		maxv = max(maxv, v)
	}
	points := make([]int, maxv+1)
	dp := make([]int, maxv+1)
	for _, v := range nums {
		points[v] += v
	}
	dp[0], dp[1] = points[0], max(points[0], points[1])
	for i := 2; i <= maxv; i++ {
		take := points[i] + dp[i-2]
		skip := dp[i-1]
		dp[i] = max(take, skip)
	}
	return dp[maxv]
}
