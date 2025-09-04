package problems

const MAXAREA int = 901

var A [30]int
var T [30]int

func getAT(grid [][]int) {
	NR, NC := len(grid), len(grid[0])
	for j := range NC {
		T[j] = 0
	}
	for i := range NR {
		A[i] = 0
		for j := range NC {
			if grid[i][j] == 1 {
				A[i] |= (1 << j)
				T[j] |= (1 << i)
			}
		}
	}
}

func getArea(r0, rN, c0, cN int) int {
	top, bottom := 30, -1
	left, right := 30, -1
	rMask := ((1 << (cN - c0 + 1)) - 1) << c0
	for r := r0; r <= rN; r++ {
		row := A[r]
		if (row & rMask) != 0 {
			top = r
			break
		}
	}
	if top == 30 {
		return MAXAREA
	}
	for r := rN; r >= top; r-- {
		row := A[r]
		if (row & rMask) != 0 {
			bottom = r
			break
		}
	}
	cMask := ((1 << (rN - r0 + 1)) - 1) << r0
	for c := c0; c <= cN; c++ {
		col := T[c]
		if (col & cMask) != 0 {
			left = c
			break
		}
	}
	for c := cN; c >= left; c-- {
		col := T[c]
		if (col & cMask) != 0 {
			right = c
			break
		}
	}
	return (bottom - top + 1) * (right - left + 1)
}

func minimumSum(grid [][]int) int {
	getAT(grid)
	NR, NC := len(grid), len(grid[0])
	res := NR * NC
	var a1, a2, a3 int
	// 3 cols ||
	for c1 := 0; c1 < NC-2; c1++ {
		a1 = getArea(0, NR-1, 0, c1)
		for c2 := c1 + 1; c2 < NC-1; c2++ {
			a2 = getArea(0, NR-1, c1+1, c2)
			a3 = getArea(0, NR-1, c2+1, NC-1)
			res = min(res, a1+a2+a3)
		}
	}
	// 3 rows =
	for r1 := 0; r1 < NR-2; r1++ {
		a1 = getArea(0, r1, 0, NC-1)
		for r2 := r1 + 1; r2 < NR-1; r2++ {
			a2 = getArea(r1+1, r2, 0, NC-1)
			a3 = getArea(r2+1, NR-1, 0, NC-1)
			res = min(res, a1+a2+a3)
		}
	}
	// T shape

	for r := 0; r < NR-1; r++ {
		for c := 0; c < NC-1; c++ {
			// bottom T
			a1 = getArea(0, r, 0, NC-1)
			a2 = getArea(r+1, NR-1, 0, c)
			a3 = getArea(r+1, NR-1, c+1, NC-1)
			res = min(res, a1+a2+a3)
			// top 1
			a1 = getArea(r+1, NR-1, 0, NC-1)
			a2 = getArea(0, r, 0, c)
			a3 = getArea(0, r, c+1, NC-1)
			res = min(res, a1+a2+a3)
			// right |-
			a1 = getArea(0, NR-1, 0, c)
			a2 = getArea(0, r, c+1, NC-1)
			a3 = getArea(r+1, NR-1, c+1, NC-1)
			res = min(res, a1+a2+a3)
			// left -|
			a1 = getArea(0, NR-1, c+1, NC-1)
			a2 = getArea(0, r, 0, c)
			a3 = getArea(r+1, NR-1, 0, c)
			res = min(res, a1+a2+a3)
		}
	}
	return res
}
