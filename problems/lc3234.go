package problems

var PZ [40000]int

func numberOfSubstrings(s string) int {
	ans := 0
	ones := 0
	PZ[0] = -1
	zl := 1
	for i, c := range s {
		if c == '0' {
			PZ[zl] = i
			zl++
		} else {
			ones++
			ans += i - PZ[zl-1]
		}
		for j := zl - 1; j >= 1; j-- {
			cur, prev := PZ[j], PZ[j-1]
			c0 := zl - j
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
