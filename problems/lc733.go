package problems

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	DIRS := [5]int{0, 1, 0, -1, 0}
	oc := image[sr][sc]
	if oc == color {
		return image
	}
	m, n := len(image), len(image[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r == m || c < 0 || c == n || image[r][c] != oc {
			return
		}
		image[r][c] = color
		for i := range 4 {
			dfs(r+DIRS[i], c+DIRS[i+1])
		}
	}
	dfs(sr, sc)
	return image
}
