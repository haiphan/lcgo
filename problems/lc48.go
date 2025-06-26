package problems

func rotate(matrix [][]int) {
	N := len(matrix)
	revRow := func(arr []int) {
		l, r := 0, N-1
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	for r := range N {
		for c := 0; c < r; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
	for r := range N {
		revRow(matrix[r])
	}
}
