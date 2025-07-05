package problems

func findLucky(arr []int) int {
	cc := make(map[int]int, len(arr))
	for _, v := range arr {
		cc[v]++
	}
	res := -1
	for k, v := range cc {
		if k == v {
			res = max(res, v)
		}
	}
	return res
}
