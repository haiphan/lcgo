package problems

func contains(w string, x byte) bool {
	for i := range len(w) {
		if w[i] == x {
			return true
		}
	}
	return false
}
func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0)
	for i := range words {
		if contains(words[i], x) {
			res = append(res, i)
		}
	}
	return res
}
