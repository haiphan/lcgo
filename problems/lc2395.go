package problems

func findSubarrays(nums []int) bool {
	N := len(nums)
	if N == 2 {
		return false
	}
	ns := make(map[int]bool)
	N1 := N - 1
	for i := range N1 {
		x := nums[i] + nums[i+1]
		if ns[x] {
			return true
		}
		ns[x] = true
	}
	return false
}
