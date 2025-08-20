package problems

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	res := 0
	for i := range m {
		for j := range n {
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				matrix[i][j] = 1 + min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1])
			}
			res += matrix[i][j]
		}
	}
	return res
}
