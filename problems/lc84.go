package problems

func largestRectangleArea(heights []int) int {
	L := len(heights)
	if L == 1 {
		return heights[0]
	}
	stack := make([]int, 0)
	res := 0
	for i := range L + 1 {
		for len(stack) > 0 && (i == L || heights[i] <= heights[stack[len(stack)-1]]) {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) > 0 {
				w = i - 1 - stack[len(stack)-1]
			}
			res = max(res, w*h)
		}
		stack = append(stack, i)
	}
	return res
}
