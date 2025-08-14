package problems

func majorityElement(nums []int) int {
	m := nums[0]
	cnt := 0
	for _, n := range nums {
		if n == m {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				cnt = 1
				m = n
			}
		}
	}
	return m
}
