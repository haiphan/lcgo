package problems

// PZ[i] = index of i-th '0' in s, PZ[0] = -1 for convenience
var PZ [40000]int

func numberOfSubstrings(s string) int {
	ans := 0
	ones := 0
	PZ[0] = -1
	// zl = number of '0' seen so far. initially 0, but we set PZ[0] = -1 for convenience, so zl starts from 1
	zl := 1
	for i, c := range s {
		if c == '0' {
			PZ[zl] = i
			zl++
		} else {
			ones++
			// add all substrings that end at i and have at least 1 '0'. the number of such substrings is i - PZ[zl-1], because we can choose any starting index from PZ[zl-1]+1 to i, and we need to have at least 1 '0' in the substring, so the starting index must be greater than PZ[zl-1].
			ans += i - PZ[zl-1]
		}
		for j := zl - 1; j >= 1; j-- {
			cur, prev := PZ[j], PZ[j-1]
			// c0 = number of '0' in the substring, c1 = number of '1' in the substring
			c0 := zl - j // To get c0, we can use zl - j, because zl is the number of '0' seen so far, and j is the index of the last '0' in the substring. So zl - j gives us the number of '0' in the substring.
			c0Sq := c0 * c0
			if c0Sq > ones {
				break
			}
			c1 := i - cur + 1 - c0
			// cand = number of substrings that end at i and have c0 '0' and c1 '1'.
			// (cur - prev) is the number of ways to choose the starting index of the substring, and max(0, c0Sq - c1) is the number of ways to choose the ending index of the substring, because we need to have at least c0Sq '1' in the substring to satisfy the condition that c0^2 > c1.
			cand := cur - prev - max(0, c0Sq-c1)
			if cand > 0 {
				ans += cand
			}
		}
	}
	return ans
}
