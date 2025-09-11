package problems

func reverseVowels(s string) string {
	n := len(s)
	arr := []byte(s)
	l, r := 0, n-1
	for l < r {
		for l < r && !isVowel(s[l]) {
			l++
		}
		for l < r && !isVowel(s[r]) {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return string(arr)
}
