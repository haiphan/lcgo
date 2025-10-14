package problems

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if (k << 1) > n {
		return false
	}
	pa := 0
	UB := n - k
	for a := 1; a <= UB; a++ {
		if pa+k == a {
			return true
		}
		b := a + k
		if b >= n {
			return false
		}
		if nums[a] <= nums[a-1] || nums[b] <= nums[b-1] {
			pa = a
		}
	}
	return false
}
