package problems

import "strings"

func minExtraChar(s string, dictionary []string) int {
	N := len(s)
	N1 := N + 1
	cache := make([]int, N, N)
	for i := range N {
		cache[i] = N1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == N {
			return 0
		}
		if cache[i] <= N {
			return cache[i]
		}
		res := 1 + dfs(i+1)
		cur := s[i:]
		for _, w := range dictionary {
			if strings.HasPrefix(cur, w) {
				res = min(res, dfs(i+len(w)))
			}
		}
		cache[i] = res
		return res
	}
	return dfs(0)
}
