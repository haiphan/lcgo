package problems

func replaceNonCoprimes(nums []int) []int {
	stack := make([]int, 0, len(nums))
	for _, x := range nums {
		for len(stack) > 0 {
			y := stack[len(stack)-1]
			g := GCD(x, y)
			if g == 1 {
				break
			}
			x = x / g * y
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, x)
	}
	return stack
}
