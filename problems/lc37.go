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
	empties := make([]int, 0, 81)
	rmNum := func(r, c, v int) {
		mask := 1 << v
		row[r] ^= mask
		col[c] ^= mask
		qua[(r/3)*3+c/3] ^= mask
		board[r][c] = '.'
	}

	var backtrack func()

	backtrack = func() {
		if done {
			return
		}
		minIdx, minCand := -1, fullmask+1
		var candVals int
		for idx, pos := range empties {
			r, c := pos/9, pos%9
			if board[r][c] != '.' {
				continue
			}
			rcq := row[r] | col[c] | qua[(r/3)*3+c/3]
			cand := (rcq ^ fullmask) & fullmask
			numCand := bits.OnesCount(uint(cand))
			if numCand < minCand {
				minCand = numCand
				minIdx = idx
				candVals = cand
				if minCand == 1 {
					break // can't get better than 1
				}
			}
		}
		if minIdx == -1 {
			done = true
			return
		}
		pos := empties[minIdx]
		r, c := pos/9, pos%9
		for cand := candVals; cand > 0; cand &= cand - 1 {
			v := bits.TrailingZeros(uint(cand))
			putNum(r, c, v, true)
			backtrack()
			if done {
				return
			}
			rmNum(r, c, v)
		}
	}
	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				empties = append(empties, i*9+j)
			} else {
				putNum(i, j, int(board[i][j]-'0'), false)
			}
		}
	}
	backtrack()
}
