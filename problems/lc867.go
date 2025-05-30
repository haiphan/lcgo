package problems

func transpose(matrix [][]int) [][]int {
	M, N := len(matrix), len(matrix[0])
	T := make([][]int, N)
	for j := range N {
		T[j] = make([]int, M)
		for i := range M {
			T[j][i] = matrix[i][j]
		}
	}
	return T
}
