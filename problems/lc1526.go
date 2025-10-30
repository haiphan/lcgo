package problems

func minNumberOperations(target []int) int {
	ans := target[0]
	for i := 1; i < len(target); i++ {
		ans += max(0, target[i]-target[i-1])
	}
	return ans
}
