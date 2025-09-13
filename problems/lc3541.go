package problems

func maxFreqSum(s string) int {
	L := len(s)
	maxv, maxc := 0, 0
	cm := [26]int{}
	for i := range L {
		ci := int(s[i] - 'a')
		isV := s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u'
		cm[ci]++
		if isV {
			maxv = max(maxv, cm[ci])
		} else {
			maxc = max(maxc, cm[ci])
		}
	}
	return maxv + maxc
}
