package problems

func sortColors(nums []int) {
	cnt := []int{0, 0, 0}
	for _, n := range nums {
		cnt[n]++
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		for cnt[j] == 0 {
			j++
		}
		cnt[j]--
		nums[i] = j
	}
}
