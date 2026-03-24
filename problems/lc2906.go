package problems

func constructProductMatrix(grid [][]int) [][]int {
	const MOD int = 12345
	m, n := len(grid), len(grid[0])
	shift := 32
	mask := (1 << shift) - 1
	pre := 1
	for i := range m {
		for j := range n {
			x := grid[i][j] % MOD
			grid[i][j] = (pre << shift) | x
			pre = (pre * x) % MOD
		}
	}
	suf := 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			pre = grid[i][j] >> shift
			x := grid[i][j] & mask
			grid[i][j] = pre * suf % MOD
			suf = (suf * x) % MOD
		}
	}
	return grid
}
