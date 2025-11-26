package problems

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	prev := make([][]int, n)
	cur := make([][]int, n)
	for i := range prev {
		prev[i] = make([]int, k)
		cur[i] = make([]int, k)
	}
	resetCur := func() {
		for i := range n {
			for j := range k {
				cur[i][j] = 0
			}
		}
	}
	for i := range m {
		resetCur()
		for j := range n {
			v := grid[i][j] % k
			if i == 0 && j == 0 {
				cur[0][v] = 1
				continue
			}
			if i > 0 {
				for r := 0; r < k; r++ {
					pc := prev[j][r]
					if pc > 0 {
						rp := (r + v) % k
						cur[j][rp] = (cur[j][rp] + pc) % MOD
					}
				}
			}
			if j > 0 {
				for r := 0; r < k; r++ {
					pc := cur[j-1][r]
					if pc > 0 {
						rp := (r + v) % k
						cur[j][rp] = (cur[j][rp] + pc) % MOD
					}
				}
			}
		}
		prev, cur = cur, prev
	}
	return prev[n-1][0]
}
