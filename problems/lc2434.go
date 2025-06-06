package problems

func robotWithString(s string) string {
	cc := make([]int, 26)
	N := len(s)
	stack := make([]byte, 0)
	res := make([]byte, 0)
	getSC := func() byte {
		for i := range 26 {
			if cc[i] > 0 {
				return byte(i + 'a')
			}
		}
		return 'z'
	}
	for _, c := range s {
		cc[int(c-'a')]++
	}
	for i := range N {
		c := s[i]
		cn := int(c - 'a')
		cc[cn]--
		stack = append(stack, c)
		sc := getSC()
		for len(stack) > 0 && stack[len(stack)-1] <= sc {
			res = append(res, stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]
		}
	}
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[0 : len(stack)-1]
	}
	return string(res)
}
