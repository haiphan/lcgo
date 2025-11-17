package problems

func kLengthApart(nums []int, k int) bool {
	prev := -k - 1
	for i, v := range nums {
		if v == 1 {
			if prev+k >= i {
				return false
			}
			prev = i
		}
	}
	return true
}
