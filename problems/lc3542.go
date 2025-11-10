package problems

func minOperations3542(nums []int) int {
	ans := 0
	stack := make([]int, 0)
	for _, x := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > x {
			stack = stack[:len(stack)-1]
		}
		if x == 0 {
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] < x {
			ans++
			stack = append(stack, x)
		}
	}
	return ans
}
