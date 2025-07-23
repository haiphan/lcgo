package problems

func maximumGain(s string, x int, y int) int {
	N := len(s)
	p1, p2 := [2]byte{'a', 'b'}, [2]byte{'b', 'a'}
	if x < y {
		x, y = y, x
		p1, p2 = p2, p1
	}
	bArr := []byte(s)
	removePa := func(p [2]byte, v int) int {
		stack := make([]byte, 0, N)
		sum := 0
		for _, b := range bArr {
			if len(stack) > 0 && stack[len(stack)-1] == p[0] && b == p[1] {
				sum += v
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, b)
			}
		}
		bArr = stack
		return sum
	}

	return removePa(p1, x) + removePa(p2, y)
}
