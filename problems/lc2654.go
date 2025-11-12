package problems

func minOperations2654(nums []int) int {
	n := len(nums)
	g := nums[0]
	c1 := 0
	for _, x := range nums {
		g = gcd(g, x)
		if x == 1 {
			c1++
		}
	}
	if g > 1 {
		return -1
	}
	if c1 > 0 {
		return n - c1
	}
	minSize := n
	for i, x := range nums {
		if minSize == 2 {
			break
		}
		for j := i + 1; j < n; j++ {
			cur := j - i + 1
			if cur == minSize {
				break
			}
			x = gcd(x, nums[j])
			if x == 1 {
				minSize = cur
				break
			}
		}
	}
	return n - 2 + minSize
}
