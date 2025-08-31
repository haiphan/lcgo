package problems

import "math/bits"

const fullmask int = 0x3fe

func solveSudoku(board [][]byte) {
	var row [9]int
	var col [9]int
	var qua [9]int
	done := false
	putNum := func(r, c, v int, assign bool) {
		mask := 1 << v
		qi := (r/3)*3 + c/3
		row[r] |= mask
		col[c] |= mask
		qua[qi] |= mask
		if assign {
			board[r][c] = byte('0' + v)
		}
	}

	rmNum := func(r, c, v int) {
		mask := 1 << v
		row[r] ^= mask
		col[c] ^= mask
		qua[(r/3)*3+c/3] ^= mask
		board[r][c] = '.'
	}

	var dfs func(i int)

	dfs = func(i int) {
		if done {
			return
		}
		if i == 81 {
			done = true
			return
		}
		r, c := i/9, i%9
		if board[r][c] == '.' {
			rcq := row[r] | col[c] | qua[(r/3)*3+c/3]
			if rcq == fullmask {
				return
			}
			cand := (rcq ^ fullmask) & fullmask
			for cand > 0 {
				v := bits.TrailingZeros(uint(cand))
				cand &= cand - 1
				putNum(r, c, v, true)
				dfs(i + 1)
				if done {
					return
				}
				rmNum(r, c, v)
			}
		} else {
			dfs(i + 1)
		}
	}
	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}
			putNum(i, j, int(board[i][j]-'0'), false)
		}
	}
	dfs(0)
}
