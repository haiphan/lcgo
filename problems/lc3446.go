package problems

import "sort"

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	for d := 1; d < n-1; d++ {
		diag := make([]int, 0, n)
		for i := 0; i+d < n; i++ {
			diag = append(diag, grid[i][i+d])
		}
		sort.Ints(diag)
		di := 0
		for i := 0; i+d < n; i++ {
			grid[i][i+d] = diag[di]
			di++
		}
	}

	for d := 2 - n; d <= 0; d++ {
		diag := make([]int, 0, n)
		for i := -d; i < n; i++ {
			diag = append(diag, grid[i][i+d])
		}
		sort.Slice(diag, func(i, j int) bool {
			return diag[i] > diag[j]
		})
		di := 0
		for i := -d; i < n; i++ {
			grid[i][i+d] = diag[di]
			di++
		}
	}

	return grid
}
