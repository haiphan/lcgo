package problems

func spiralOrder(matrix [][]int) []int {
	M, N := len(matrix), len(matrix[0])
	top, bottom := 0, M-1
	left, right := 0, N-1
	res := make([]int, 0, M*N)
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- {
				res = append(res, matrix[r][left])
			}
			left++
		}
	}
	return res
}
