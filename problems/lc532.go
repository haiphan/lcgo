package problems

func findDup(nums []int) int {
	N := len(nums)
	nc := make(map[int]bool, N)
	done := make(map[int]bool, N)
	res := 0
	for _, x := range nums {
		if done[x] {
			continue
		}
		if nc[x] {
			done[x] = true
			res++
		} else {
			nc[x] = true
		}
	}
	return res
}

func findPairs(nums []int, k int) int {
	if k == 0 {
		return findDup(nums)
	}
	N := len(nums)
	nc := make(map[int]bool, N)
	for _, x := range nums {
		nc[x] = true
	}
	res := 0
	for x := range nc {
		_, has := nc[x+k]
		if has {
			res++
		}
	}
	return res
}
