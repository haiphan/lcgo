package problems

func findLHS(nums []int) int {
	N := len(nums)
	cm := make(map[int]int, N)
	for _, x := range nums {
		cm[x]++
	}
	res := 0
	for k, v := range cm {
		u, has := cm[k+1]
		if has {
			res = max(res, v+u)
		}
	}
	return res
}
