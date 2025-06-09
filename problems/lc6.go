package problems

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	N := len(s)
	cycleLen := 2*numRows - 2
	res := make([]byte, 0, N)
	for r := range numRows {
		for i := r; i < N; i += cycleLen {
			res = append(res, s[i])
			// For middle rows, add the character from the upward stroke
			if r != 0 && r != numRows-1 {
				j := i + cycleLen - 2*r
				if j < N {
					res = append(res, s[j])
				}
			}
		}
	}
	return string(res)
}
