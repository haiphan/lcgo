package problems

func countHillValley(nums []int) int {
	UB := len(nums) - 1
	l, cnt := 0, 0
	for i := 1; i < UB; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		if (nums[i]-nums[l])*(nums[i]-nums[i+1]) > 0 {
			cnt++
		}
		l = i
	}
	return cnt
}
