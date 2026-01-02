package problems

func repeatedNTimes(nums []int) int {
	cm := make(map[int]bool, len(nums))
	for _, x := range nums {
		if cm[x] {
			return x
		}
		cm[x] = true
	}
	return 0
}
