package problems

func hasSameDigits(s string) bool {
	n := len(s)
	arr := make([]int, n)
	for i := range n {
		arr[i] = int(s[i] - '0')
	}
	for ; n > 2; n-- {
		for i := 0; i+1 < n; i++ {
			arr[i] = (arr[i] + arr[i+1]) % 10
		}
	}
	return arr[0] == arr[1]
}
