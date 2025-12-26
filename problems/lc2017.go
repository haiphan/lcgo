package problems

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	topSum := 0
	for _, x := range grid[0] {
		topSum += x
	}
	bottomSum := 0
	ans := topSum
	for i := range n {
		topSum -= grid[0][i]
		ans = min(ans, max(topSum, bottomSum))
		bottomSum += grid[1][i]
	}
	return int64(ans)
}
