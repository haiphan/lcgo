package problems

func makeFancyString(s string) string {
	N := len(s)
	res := make([]byte, 0)
	var prev byte
	cnt := 0
	for i := range N {
		c := s[i]
		if c == prev {
			if cnt > 1 {
				continue
			}
			cnt++
		} else {
			prev = c
			cnt = 1
		}
		res = append(res, c)
	}
	return string(res)
}
