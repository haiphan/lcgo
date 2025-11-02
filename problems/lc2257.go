package problems

const XV byte = 1
const EV byte = 2

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	var dp = make([]byte, m*n)
	getIndex := func(r, c int) int {
		return r*n + c
	}
	extendG := func(r, c int) bool {
		i := getIndex(r, c)
		if dp[i] == XV {
			return true
		}
		dp[i] = EV
		return false
	}
	for _, w := range walls {
		dp[getIndex(w[0], w[1])] = XV
	}
	for _, g := range guards {
		dp[getIndex(g[0], g[1])] = XV
	}
	for _, g := range guards {
		r, c := g[0], g[1]
		for nr := r - 1; nr >= 0; nr-- {
			if extendG(nr, c) {
				break
			}
		}
		for nr := r + 1; nr < m; nr++ {
			if extendG(nr, c) {
				break
			}
		}
		for nc := c - 1; nc >= 0; nc-- {
			if extendG(r, nc) {
				break
			}
		}
		for nc := c + 1; nc < n; nc++ {
			if extendG(r, nc) {
				break
			}
		}
	}
	ans := 0
	for _, x := range dp {
		if x == 0 {
			ans++
		}
	}
	return ans
}
