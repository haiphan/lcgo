package problems

func findErrorNums(nums []int) []int {
	n := len(nums)
	seen := make([]bool, n+1)
	ans := []int{0, 0}
	sum := 0
	// D = b - a, b = D + a
	for _, x := range nums {
		if seen[x] {
			ans[0] = x
		}
		sum += x
		seen[x] = true
	}
	exp := (1 + n) * n / 2
	ans[1] = exp - sum + ans[0]
	return ans
}
