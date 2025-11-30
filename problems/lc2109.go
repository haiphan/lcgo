package problems

func addSpaces(s string, spaces []int) string {
	si := 0
	n, m := len(s), len(spaces)
	res := make([]byte, 0, n+m)
	spaces = append(spaces, -1)
	for i := range s {
		if i == spaces[si] {
			si++
			res = append(res, ' ')
		}
		res = append(res, s[i])
	}
	return string(res)
}
