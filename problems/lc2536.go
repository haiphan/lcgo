package problems

func rangeAddQueries(n int, queries [][]int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		ans[r1][c1]++
		if r2+1 < n {
			ans[r2+1][c1]--
			if c2+1 < n {
				ans[r2+1][c2+1]++
			}
		}
		if c2+1 < n {
			ans[r1][c2+1]--
		}
	}
	for r := range n {
		for c := 1; c < n; c++ {
			ans[r][c] += ans[r][c-1]
		}
	}
	for r := 1; r < n; r++ {
		for c := range n {
			ans[r][c] += ans[r-1][c]
		}
	}
	return ans
}
