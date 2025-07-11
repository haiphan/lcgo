package problems

func sortArray(nums []int) []int {
	cm := make(map[int]int)
	minN, maxN := nums[0], nums[0]
	for _, n := range nums {
		minN, maxN = min(minN, n), max(maxN, n)
		cm[n]++
	}
	i := 0
	for j := minN; j <= maxN; j++ {
		if k, has := cm[j]; has {
			for ; k > 0; k-- {
				nums[i] = j
				i++
			}
		}
	}
	return nums
}
