package problems

func isValidSudoku(board [][]byte) bool {
	var rows [9]int
	var cols [9]int
	var quads [9]int
	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}
			mask := 1 << int(board[i][j]-'0')
			qi := (i/3)*3 + (j / 3)
			if (rows[i]&mask) == 0 && (cols[j]&mask) == 0 && (quads[qi]&mask) == 0 {
				rows[i] |= mask
				cols[j] |= mask
				quads[qi] |= mask
			} else {
				return false
			}
		}
	}
	return true
}
