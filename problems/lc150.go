package problems

import "strconv"

func doop(a, b int, c string) int {
	switch c {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
func evalRPN(tokens []string) int {
	N := len(tokens)
	stack := make([]int, 0, N)
	syms := make(map[string]bool, N)
	syms["+"] = true
	syms["-"] = true
	syms["*"] = true
	syms["/"] = true
	for _, s := range tokens {
		if syms[s] {
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, doop(a, b, s))
		} else {
			si, _ := strconv.Atoi(s)
			stack = append(stack, si)
		}
	}
	return stack[0]
}
