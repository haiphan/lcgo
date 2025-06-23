package problems

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}
	tc := make([]int, n+1)
	hasTrust := make(map[int]bool, n)
	candidates := make(map[int]bool, n)
	n1 := n - 1
	for _, t := range trust {
		tc[t[1]]++
		hasTrust[t[0]] = true
		if tc[t[1]] == n1 {
			candidates[t[1]] = true
		}
	}
	for k := range candidates {
		if !hasTrust[k] {
			return k
		}
	}
	return -1
}
