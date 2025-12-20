package problems

func minDeletionSize(strs []string) int {
	ans := 0
	n := len(strs[0])
	for c := range n {
		for r := 1; r < len(strs); r++ {
			if strs[r][c] < strs[r-1][c] {
				ans++
				break
			}
		}
	}
	return ans
}
