package problems

func getRow(rowIndex int) []int {
	n := rowIndex
	line := make([]int, n+1)
	line[0] = 1
	line[n] = 1
	half := n >> 1

	for k := 1; k <= half; k++ {
		line[k] = line[k-1] * (n - k + 1) / k
		line[n-k] = line[k]
	}

	return line
}
