package problems

const BS int = 8

func numRookCaptures(board [][]byte) int {
	r, c := 0, 0
	ans := 0
	for i := range BS {
		for j := range BS {
			if board[i][j] == 'R' {
				r, c = i, j
				break
			}
		}
	}
	shouldStop := func(i, j int) bool {
		if i < 0 || i == BS || j < 0 || j == BS {
			return true
		}
		if board[i][j] == '.' {
			return false
		}
		if board[i][j] == 'p' {
			ans++
		}
		return true
	}

	for di := range 4 {
		dr, dc := gridDirs[di], gridDirs[di+1]
		nr, nc := r+dr, c+dc
		for !shouldStop(nr, nc) {
			nr, nc = nr+dr, nc+dc
		}
	}
	return ans
}
