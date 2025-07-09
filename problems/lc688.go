package problems

var DIRS [8][2]int = [8][2]int{
	{-2, -1},
	{-2, 1},
	{2, -1},
	{2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
}

func getBoard(n int) [][]float64 {
	b := make([][]float64, n)
	for i := range n {
		b[i] = make([]float64, n)
	}
	return b
}

func knightProbability(n int, k int, row int, column int) float64 {
	cur, next := getBoard(n), getBoard(n)
	cur[row][column] = 1.0
	for k > 0 {
		for r := range n {
			for c := range n {
				if cur[r][c] == 0 {
					continue
				}
				for _, d := range DIRS {
					nr, nc := r+d[0], c+d[1]
					if nr < 0 || nr >= n || nc < 0 || nc >= n {
						continue
					}
					next[nr][nc] += (cur[r][c] * 0.125)
				}
				cur[r][c] = 0
			}
		}
		cur, next = next, cur
		k--
	}
	p := 0.0
	for r := range n {
		for c := range n {
			p += cur[r][c]
		}
	}
	return p
}
