package problems

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i, end := 0, n-2
	for i <= end {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == n-1
}
