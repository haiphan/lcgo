package problems

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, n := range asteroids {
		if n > 0 {
			stack = append(stack, n)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 {
				last := stack[len(stack)-1]
				sum := last + n
				if sum == 0 {
					stack = stack[:len(stack)-1]
					n = 0
					break
				} else if sum > 0 {
					n = 0
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if n != 0 {
				stack = append(stack, n)
			}
		}
	}
	return stack
}
