package problems

import "slices"

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	top, bottom, left, right := m, -1, n, -1
	colHasOne := func(j int) bool {
		for i := range m {
			if grid[i][j] == 1 {
				return true
			}
		}
		return false
	}
	for i := range m {
		if slices.Contains(grid[i], 1) {
			top = i
			break
		}
	}
	for i := m - 1; i >= 0; i-- {
		if slices.Contains(grid[i], 1) {
			bottom = i
			break
		}
	}
	for j := range n {
		if colHasOne(j) {
			left = j
			break
		}
	}
	for j := n - 1; j >= 0; j-- {
		if colHasOne(j) {
			right = j
			break
		}
	}
	return (bottom - top + 1) * (right - left + 1)
}
