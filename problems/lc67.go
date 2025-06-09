package problems

import "strconv"

func addBinary(a string, b string) string {
	AL, BL := len(a), len(b)
	ML := max(AL, BL)
	res := ""
	c := 0
	for i := range ML {
		ia, ib := AL-1-i, BL-1-i
		va, vb := 0, 0
		if ia >= 0 {
			va = int(a[ia] - '0')
		}
		if ib >= 0 {
			vb = int(b[ib] - '0')
		}
		d := va + vb + c
		c = 0
		if d > 1 {
			c = 1
			d = d % 2
		}
		res = strconv.Itoa(d) + res
	}
	if c == 1 {
		res = "1" + res
	}
	return res
}
