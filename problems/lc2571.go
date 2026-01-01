package problems

func minOperations2571(n int) int {
	ans := 0
	for n > 0 {
		if (n & 3) == 3 {
			n++
			ans++
		} else {
			ans += (n & 1)
			n >>= 1
		}
	}
	return ans
}
