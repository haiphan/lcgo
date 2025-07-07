package problems

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	N := len(nums)
	for i := range N {
		for j := i + indexDifference; j < N; j++ {
			if abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
