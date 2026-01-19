package problems

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	ps := make([][]int, m+1)
	for i := range ps {
		ps[i] = make([]int, n+1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			// Formula: Current + Top + Left - TopLeftDiagonal
			ps[r+1][c+1] = mat[r][c] + ps[r][c+1] + ps[r+1][c] - ps[r][c]
		}
	}
	getSum := func(i, j, l int) int {
		// Convert coordinates to 1-based for the prefix table
		// Top-left: (i, j) -> (r1, c1)
		// Bottom-right: (i+l-1, j+l-1) -> (r2, c2)
		r1, c1 := i+1, j+1
		r2, c2 := i+l, j+l

		// Result = BottomRight - Top - Left + TopLeftOverlap
		return ps[r2][c2] - ps[r1-1][c2] - ps[r2][c1-1] + ps[r1-1][c1-1]
	}
	maxl := 0
	for i := range m {
		if i+maxl >= m {
			break
		}
		for j := 0; j < n-maxl; j++ {
			if j+maxl >= n {
				break
			}
			for i+maxl < m && j+maxl < n && getSum(i, j, maxl+1) <= threshold {
				maxl++
			}
		}
	}
	return maxl
}
