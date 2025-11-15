package problems

func numberOfSubstrings(s string) int {
	ans := 0
	ones := 0
	pz := []int{-1}
	for i, c := range s {
		if c == '0' {
			pz = append(pz, i)
		} else {
			ones++
			ans += i - pz[len(pz)-1]
		}
		lz := len(pz)
		for j := lz - 1; j >= 1; j-- {
			cur, prev := pz[j], pz[j-1]
			c0 := lz - j
			c0Sq := c0 * c0
			if c0Sq > ones {
				break
			}
			c1 := i - cur + 1 - c0
			cand := cur - prev - max(0, c0Sq-c1)
			if cand > 0 {
				ans += cand
			}
		}
	}
	return ans
}
