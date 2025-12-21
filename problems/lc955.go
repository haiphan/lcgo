package problems

func minDeletionSize955(strs []string) int {
	n := len(strs)
	m := len(strs[0])
	resolved := make([]bool, n)
	ans := 0
	for j := range m {
		// column j can be kept
		take := true
		for i := 0; i < n-1; i++ {
			if resolved[i] {
				continue
			}
			if strs[i][j] > strs[i+1][j] {
				take = false
				break
			}
		}
		if take {
			// row i is resolved if strs[i][j] < strs[i+1][j]
			for i := 0; i < n-1; i++ {
				if !resolved[i] && strs[i][j] < strs[i+1][j] {
					resolved[i] = true
				}
			}
		} else {
			ans++
		}
	}
	return ans
}
