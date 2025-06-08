package problems

func lexicalOrder(n int) []int {
	res := make([]int, n)
	i, cur := 0, 1
	for i < n {
		res[i] = cur
		i++
		c10 := cur * 10
		if c10 <= n {
			cur = c10
		} else {
			for cur == n || cur%10 == 9 {
				cur /= 10
			}
			cur++
		}
	}
	return res
}
