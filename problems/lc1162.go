package problems

const FARDIST int = 1000

func maxDistance1162(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return -1
	}
	n1 := n - 1
	TK := n << 1
	for i := range n {
		for j := range n {
			if grid[i][j] == 1 {
				continue
			}
			top, left := FARDIST, FARDIST
			if i != 0 {
				top = grid[i-1][j] >> 1
			}
			if j != 0 {
				left = grid[i][j-1] >> 1
			}
			grid[i][j] = (min(top, left) + 1) << 1
		}
	}
	ans := -1
	for i := n1; i >= 0; i-- {
		for j := n1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			bottom, right := FARDIST, FARDIST
			if i != n1 {
				bottom = grid[i+1][j] >> 1
			}
			if j != n1 {
				right = grid[i][j+1] >> 1
			}
			cur := grid[i][j] >> 1
			cur = min(cur, bottom+1, right+1)
			grid[i][j] = cur << 1
			if ans < cur && cur < TK {
				ans = cur
			}
		}
	}
	return ans
}
