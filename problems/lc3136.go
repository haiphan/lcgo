package problems

import "unicode"

func isValidString(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasv, hasc := false, false
	for _, c := range word {
		if unicode.IsLetter(c) {
			c := unicode.ToLower(c)
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				hasv = true
			} else {
				hasc = true
			}
		} else if !unicode.IsDigit(c) {
			return false
		}
	}
	return hasv && hasc
}
