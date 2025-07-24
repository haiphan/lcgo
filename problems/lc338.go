package problems

func countBits(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i>>1] + i&1
	}
	return arr
}
