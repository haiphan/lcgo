package problems

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	M := len(meetings)
	ans := make([]int, 0, n)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	par := make([]int, n)
	for i := range par {
		par[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if par[x] != x {
			par[x] = find(par[x])
		}
		return par[x]
	}

	union := func(a, b int) {
		par[find(b)] = find(a)
	}

	resetNode := func(r0, x int) {
		if find(x) != r0 {
			par[x] = x
		}
	}

	union(0, firstPerson)
	i := 0
	for i < M {
		time := meetings[i][2]
		l := i
		for i < M && meetings[i][2] == time {
			union(meetings[i][0], meetings[i][1])
			i++
		}
		r0 := find(0)
		for ; l < i; l++ {
			resetNode(r0, meetings[l][0])
			resetNode(r0, meetings[l][1])
		}
	}
	r0 := find(0)
	for i := range n {
		if find(i) == r0 {
			ans = append(ans, i)
		}
	}
	return ans
}
