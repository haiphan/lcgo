package problems

var HistStack [200]int

func maxHistArea(hh []int) int {
	n := len(hh)
	// Monotonic stack, increasing
	st := HistStack[:0]
	maxArea := 0

	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = hh[i]
		}
		for len(st) > 0 && h < hh[st[len(st)-1]] {
			height := hh[st[len(st)-1]]
			st = st[:len(st)-1]
			// width calculation. If stack is empty, means all previous are >= height
			width := i
			if len(st) > 0 {
				// width, start from st[len(st)-1]+1
				// ends at i-1
				width = i - st[len(st)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		st = append(st, i)
	}
	return maxArea
}
func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	hist := make([]int, n)
	var ans int

	for i := range m {
		for j := range n {
			if matrix[i][j] == '1' {
				hist[j]++
			} else {
				hist[j] = 0
			}
		}
		ans = max(ans, maxHistArea(hist))
	}
	return ans
}
