package problems

func isZeroArray(nums []int, queries [][]int) bool {
	N := len(nums)
	prefix := make([]int, N+1)
	for _, q := range queries {
		l, r := q[0], q[1]
		prefix[l]++
		prefix[r+1]--
	}
	cur := 0
	for i, num := range nums {
		cur += prefix[i]
		if num > cur {
			return false
		}
	}
	return true
}
