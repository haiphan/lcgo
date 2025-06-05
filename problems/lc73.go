package problems

func setZeroes(matrix [][]int) {
	M, N := len(matrix), len(matrix[0])
	fr, fc := false, false
	for i := range N {
		if matrix[0][i] == 0 {
			fr = true
			break
		}
	}
	for i := range M {
		if matrix[i][0] == 0 {
			fc = true
			break
		}
	}
	for r := 1; r < M; r++ {
		for c := 1; c < N; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}
	for r := 1; r < M; r++ {
		for c := 1; c < N; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}
	if fr {
		for i := range N {
			matrix[0][i] = 0
		}
	}
	if fc {
		for i := range M {
			matrix[i][0] = 0
		}
	}
}
