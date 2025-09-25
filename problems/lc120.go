package problems

func minimumTotal(triangle [][]int) int {
	M := len(triangle)
	for i := M - 2; i >= 0; i-- {
		arr := triangle[i]
		prev := triangle[i+1]
		for j := range arr {
			arr[j] += min(prev[j], prev[j+1])
		}
	}
	return triangle[0][0]
}
