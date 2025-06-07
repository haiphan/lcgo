package problems

func clearStars(s string) string {
	N := len(s)
	cm := [26][]int{}
	ba := []byte(s)
	for i, b := range ba {
		if b == '*' {
			for j := range 26 {
				if len(cm[j]) > 0 {
					ri := cm[j][len(cm[j])-1]
					cm[j] = cm[j][:len(cm[j])-1]
					ba[ri] = '*'
					break
				}
			}
		} else {
			ci := b - 'a'
			cm[ci] = append(cm[ci], i)
		}
	}
	res := make([]byte, 0, N)
	for _, b := range ba {
		if b != '*' {
			res = append(res, b)
		}
	}
	return string(res)
}
