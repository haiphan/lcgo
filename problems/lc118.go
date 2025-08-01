package problems

func getPascalRow(n int) []int {
	if n < 0 {
		return []int{}
	}

	row := make([]int, n+1)
	row[0] = 1

	for k := 1; k <= n; k++ {
		row[k] = row[k-1] * (n - k + 1) / k
	}

	return row
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range numRows {
		res[i] = getPascalRow(i)
	}
	return res
}
