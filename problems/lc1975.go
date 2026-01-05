package problems

func maxMatrixSum(matrix [][]int) int64 {
	var BIG int = 1e5
	minv := BIG + 1
	n := len(matrix)
	neg := 0
	posSum := 0
	hasZ := false
	for i := range n {
		for j := range n {
			v := matrix[i][j]
			pv := abs(v)
			if v == 0 {
				hasZ = true
			}
			minv = min(minv, pv)
			if v < 0 {
				neg++
			}
			posSum += pv
		}
	}
	if hasZ || (neg%2) == 0 {
		return int64(posSum)
	}
	return int64(posSum - 2*minv)
}
