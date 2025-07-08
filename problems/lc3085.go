package problems

func minimumDeletions(word string, k int) int {
	cm := [26]int{}
	N := len(word)
	for i := range N {
		cm[word[i]-'a']++
	}
	freqs := make([]int, 0, 26)
	fset := make(map[int]bool, 26)
	for _, f := range cm {
		if f > 0 {
			fset[f] = true
			freqs = append(freqs, f)
		}
	}
	res := N
	for minF := range fset {
		cur := 0
		for _, f := range freqs {
			if f < minF {
				cur += f
			} else if f > minF+k {
				cur += f - minF - k
			}
		}
		res = min(res, cur)
	}
	return res
}
