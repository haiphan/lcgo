package problems

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i, end := 0, n-2
	for i <= end {
		i += 1 + bits[i]
	}
	return i == n-1
}
