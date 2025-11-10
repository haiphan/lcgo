package problems

func minOperations3542(nums []int) int {
	ans := 0
	stack := make([]int, 0)
	for _, x := range nums {
		if x == 0 {
			stack = stack[:0]
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > x {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 || stack[len(stack)-1] < x {
			ans++
			stack = append(stack, x)
		}
	}
	return ans
}
