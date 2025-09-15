package problems

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	bb := make([]bool, 26)
	for _, c := range brokenLetters {
		bb[c-'a'] = true
	}
	ww := strings.Split(text, " ")
	cnt := 0
	for _, w := range ww {
		x := 1
		for _, c := range w {
			if bb[c-'a'] {
				x = 0
				break
			}
		}
		cnt += x
	}
	return cnt
}
