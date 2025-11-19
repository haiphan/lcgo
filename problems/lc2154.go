package problems

func findFinalValue(nums []int, original int) int {
	ns := make(map[int]bool)
	for _, x := range nums {
		ns[x] = true
	}
	for ns[original] {
		original *= 2
	}
	return original
}
