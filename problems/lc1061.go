package problems

func getPar() []int {
	par := make([]int, 26)
	for i := range 26 {
		par[i] = i
	}
	return par
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	par := getPar()
	N, B := len(s1), len(baseStr)
	var find func(x int) int
	find = func(x int) int {
		if par[x] == x {
			return x
		}
		par[x] = find(par[x])
		return par[x]
	}
	union := func(x, y int) {
		small, big := find(x), find(y)
		if small == big {
			return
		}
		if small > big {
			small, big = big, small
		}
		par[big] = small
	}

	for i := range N {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}
	arr := make([]byte, B)
	for i := range B {
		arr[i] = byte(find(int(baseStr[i]-'a'))) + 'a'
	}
	return string(arr)
}
