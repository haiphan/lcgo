package problems

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	cur := 0
	for i, x := range nums {
		cur = ((cur << 1) + x) % 5
		ans[i] = cur == 0
	}
	return ans
}
