package problems

func findSmallestInteger(nums []int, value int) int {
	n := len(nums)
	cc := make([]int, value)
	for _, x := range nums {
		cc[(x%value+value)%value]++
	}
	for i := range n {
		ii := i % value
		if cc[ii] == 0 {
			return i
		}
		cc[ii]--
	}
	return n
}
