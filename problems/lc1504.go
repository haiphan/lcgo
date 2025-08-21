package problems

func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	h := make([]int, n, n)
	stack := make([]int, 0, n)
	res := 0
	countRow := func(r int) int {
		stack = stack[:0]
		cnt := 0
		rSum := mat[r]
		for i := range n {
			for len(stack) > 0 && h[stack[len(stack)-1]] >= h[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				rSum[i] = h[i] * (i + 1)
			} else {
				p := stack[len(stack)-1]
				rSum[i] = rSum[p] + h[i]*(i-p)
			}
			cnt += rSum[i]
			stack = append(stack, i)
		}
		return cnt
	}
	for i := range m {
		for j := range n {
			h[j]++
			if mat[i][j] == 0 {
				h[j] = 0
			}
		}
		res += countRow(i)
	}
	return res
}
