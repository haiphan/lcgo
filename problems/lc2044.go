package problems

func countMaxOrSubsets(nums []int) int {
	N := len(nums)
	maxv := 0
	for _, x := range nums {
		maxv |= x
	}
	var dfs func(i, cur int) int
	dfs = func(i, cur int) int {
		if cur == maxv {
			return 1 << (N - i)
		}
		if i == N {
			return 0
		}
		return dfs(i+1, cur|nums[i]) + dfs(i+1, cur)
	}
	return dfs(0, 0)
}
