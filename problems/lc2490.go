package problems

func isCircularSentence(sentence string) bool {
	n := len(sentence)
	if sentence[n-1] != sentence[0] {
		return false
	}
	i := 0
	for i < n {
		for i < n && sentence[i] != ' ' {
			i++
		}
		if i < n {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
			i++
		}
	}
	return true
}
