package problems

func getSneakyNumbers(nums []int) []int {
	ns := make(map[int]bool)
	res := make([]int, 0, 2)
	for _, x := range nums {
		if ns[x] {
			res = append(res, x)
		} else {
			ns[x] = true
		}
	}
	return res
}
