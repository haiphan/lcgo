package problems

import "math"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	know := make([][]bool, m)
	need := make([]bool, m)
	for i, langs := range languages {
		know[i] = make([]bool, n+1)
		for _, l := range langs {
			know[i][l] = true
		}
	}
	canTalk := func(a, b int) bool {
		for i := 1; i <= n; i++ {
			if know[a][i] && know[b][i] {
				return true
			}
		}
		return false
	}
	for _, fs := range friendships {
		f1, f2 := fs[0]-1, fs[1]-1
		if canTalk(f1, f2) {
			continue
		}
		need[f1] = true
		need[f2] = true
	}
	minLearn := math.MaxInt
	for i := 1; i <= n; i++ {
		cnt := 0
		for j := range m {
			if need[j] && !know[j][i] {
				cnt++
				if cnt >= minLearn {
					break
				}
			}
		}
		minLearn = min(minLearn, cnt)
	}
	return minLearn
}
