package problems

func containsDuplicate(nums []int) bool {
	nm := map[int]bool{}
	for _, value := range nums {
		if nm[value] {
			return true
		}
		nm[value] = true
	}
	return false
}
