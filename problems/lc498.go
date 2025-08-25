package problems

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	r, c := 0, 0
	up := true
	L := m * n
	res := make([]int, L)
	for i := range L {
		res[i] = mat[r][c]
		if up {
			if c+1 == n {
				r++
				up = false
			} else if r == 0 {
				c++
				up = false
			} else {
				r--
				c++
			}
		} else {
			if r+1 == m {
				c++
				up = true
			} else if c == 0 {
				r++
				up = true
			} else {
				r++
				c--
			}
		}
	}
	return res
}
