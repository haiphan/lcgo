package problems

func findKDistantIndices(nums []int, key int, k int) []int {
	N, end := len(nums), 0
	res := make([]int, 0, N)
	for i := range N {
		if nums[i] == key {
			start := max(end, i-k)
			end = min(N-1, i+k) + 1
			for j := start; j < end; j++ {
				res = append(res, j)
			}
		}
	}
	return res
}
