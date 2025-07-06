package problems

func countPairs(nums []int, k int) int {
	N := len(nums)
	cnt := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}
