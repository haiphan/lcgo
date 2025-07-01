package problems

func possibleStringCount(word string) int {
	N := len(word)
	N1 := N - 1
	cnt := 1
	for i := range N1 {
		if word[i] == word[i+1] {
			cnt++
		}
	}
	return cnt
}
