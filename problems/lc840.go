package problems

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	isMS := func(r, c int) bool {
		var cc [10]int
		for i := r; i < r+3; i++ {
			for j := c; j < c+3; j++ {
				if grid[i][j] <= 0 || grid[i][j] > 9 || cc[grid[i][j]] > 0 {
					return false
				}
				cc[grid[i][j]]++
			}
		}
		sum := grid[r][c] + grid[r][c+1] + grid[r][c+2]
		for i := r; i < r+3; i++ {
			if sum != grid[i][c]+grid[i][c+1]+grid[i][c+2] {
				return false
			}
		}
		for i := c; i < c+3; i++ {
			if sum != grid[r][i]+grid[r+1][i]+grid[r+2][i] {
				return false
			}
		}
		return (sum == grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2]) && (sum == grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c])
	}
	ans := 0
	for i := range m - 2 {
		for j := range n - 2 {
			if isMS(i, j) {
				ans++
			}
		}
	}
	return ans
}
