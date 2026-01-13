package problems

func strStr(haystack string, needle string) int {
	// return strings.Index(haystack, needle)
	N := len(needle)
	M := len(haystack)
	for i := 0; i < M-N+1; i++ {
		if haystack[i] == needle[0] {
			match := true
			for j := 0; j < N; j++ {
				if needle[j] != haystack[i+j] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}
