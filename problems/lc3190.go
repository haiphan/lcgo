package problems

func minimumOperations(nums []int) int {
	ans := 0
	for _, x := range nums {
		ans += min(1, x%3)
	}
	return ans
}
