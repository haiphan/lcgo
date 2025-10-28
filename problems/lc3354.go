package problems

func countValidSelections(nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, v := range nums {
		pre[i+1] = pre[i] + v
	}
	ans := 0
	for i, v := range nums {
		if v == 0 {
			cur := abs(pre[n] - (pre[i] << 1))
			switch cur {
			case 0:
				ans += 2
			case 1:
				ans += 1
			}
		}
	}
	return ans
}
