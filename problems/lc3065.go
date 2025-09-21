package problems

func minOperations1(nums []int, k int) int {
	c := 0
	for _, x := range nums {
		if x < k {
			c++
		}
	}
	return c
}
