package problems

const MAX_CITY = 100000

var par [MAX_CITY + 1]int

func processQueries(c int, connections [][]int, queries [][]int) []int {
	for i := 1; i <= c; i++ {
		par[i] = i
	}
	var find func(u int) int
	var union func(u, v int) bool
	find = func(u int) int {
		if par[u] == u {
			return u
		}
		par[u] = find(par[u])
		return par[u]
	}
	union = func(u, v int) bool {
		pu, pv := find(u), find(v)
		if pu == pv {
			return false
		}
		par[pu] = pv
		return true
	}
	for _, conn := range connections {
		union(conn[0], conn[1])
	}
	offCnt := make([]int, c+1)
	minOn := make([]int, c+1)
	for i := 0; i <= c; i++ {
		minOn[i] = -1
	}
	ansLen := 0
	for _, q := range queries {
		t, x := q[0], q[1]
		if t == 1 {
			ansLen++
		} else {
			offCnt[x]++
		}
	}
	ans := make([]int, ansLen)
	for i := 1; i <= c; i++ {
		if offCnt[i] == 0 {
			p := find(i)
			if minOn[p] == -1 || minOn[p] > i {
				minOn[p] = i
			}
		}
	}
	j := ansLen - 1

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		t, x := q[0], q[1]
		p := find(x)
		if t == 1 {
			v := x
			if offCnt[x] > 0 {
				v = minOn[p]
			}
			ans[j] = v
			j--
		} else {
			offCnt[x]--
			if offCnt[x] == 0 {
				if minOn[p] == -1 || minOn[p] > x {
					minOn[p] = x
				}
			}
		}
	}
	return ans
}
