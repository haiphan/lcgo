package problems

func largestMagicSquare(grid [][]int) int {
	NR := len(grid)
	NC := len(grid[0])
	// colPre[i][j] is the sum of column j from row 0 to row i
	// rowPre[i][j] is the sum of row i from column 0 to column j
	// d1Pre[i][j] is the sum of diagonal from (0, j-i) to (i, j)
	// d2Pre[i][j] is the sum of diagonal from (0, j+i) to (i, j)
	colPre := make([][]int, NR)
	rowPre := make([][]int, NR)
	d1Pre := make([][]int, NR)
	d2Pre := make([][]int, NR)
	for i := range NR {
		colPre[i] = make([]int, NC)
		rowPre[i] = make([]int, NC)
		d1Pre[i] = make([]int, NC)
		d2Pre[i] = make([]int, NC)
	}
	ans := 1
	// Helper to check if all sums in a range match target
	checkRange := func(pre [][]int, fixedIdx, rangeStart, rangeEnd, targetSum int, isRow bool) bool {
		if isRow {
			// Check each row in [rangeStart, rangeEnd]
			for idx := rangeStart; idx <= rangeEnd; idx++ {
				s := pre[idx][fixedIdx]
				if rangeStart > 0 {
					s -= pre[idx][rangeStart-1]
				}
				if s != targetSum {
					return false
				}
			}
		} else {
			// Check each column in [rangeStart, rangeEnd]
			for idx := rangeStart; idx <= rangeEnd; idx++ {
				s := pre[fixedIdx][idx]
				if rangeStart > 0 {
					s -= pre[rangeStart-1][idx]
				}
				if s != targetSum {
					return false
				}
			}
		}
		return true
	}
	for i := range NR {
		for j := range NC {
			v := grid[i][j]
			colPre[i][j], rowPre[i][j], d1Pre[i][j], d2Pre[i][j] = v, v, v, v
			if i > 0 {
				colPre[i][j] += colPre[i-1][j]
			}
			if j > 0 {
				rowPre[i][j] += rowPre[i][j-1]
			}
			if i > 0 && j > 0 {
				d1Pre[i][j] += d1Pre[i-1][j-1]
			}
			if i > 0 && j < NC-1 {
				d2Pre[i][j] += d2Pre[i-1][j+1]
			}
			maxK := min(i, j)
			for k := maxK; k >= ans; k-- {
				sum := colPre[i][j]
				valid := true
				if i-k > 0 {
					sum -= colPre[i-k-1][j]
				}
				if !checkRange(colPre, i, j-k, j, sum, false) {
					valid = false
				}
				if !valid {
					continue
				}
				if !checkRange(rowPre, j, i-k, i, sum, true) {
					valid = false
				}
				dSum := d1Pre[i][j]
				if i-k > 0 && j-k > 0 {
					dSum -= d1Pre[i-k-1][j-k-1]
				}
				if dSum != sum {
					continue
				}
				dSum = d2Pre[i][j-k]
				if i-k > 0 && j < NC-1 {
					dSum -= d2Pre[i-k-1][j+1]
				}
				if dSum != sum {
					continue
				}
				// Add 1 to k to convert from index to size
				ans = k + 1
				break
			}
		}
	}
	return ans
}
