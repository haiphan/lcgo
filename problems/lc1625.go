package problems

func findSmallest(a, k int) int {
	g := gcd(k, 10)
	return a % g
}

func addChar(arr []byte, k, start int) {
	x := int(arr[start] - '0')
	y := findSmallest(x, k)
	if x == y {
		return
	}
	d := (y - x + 10) % 10
	for i := start; i < len(arr); i += 2 {
		u := int(arr[i] - '0')
		v := (u + d) % 10
		arr[i] = byte('0' + v)
	}
}

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	g := gcd(n, b)
	oddB := (b & 1) == 1
	ans := s
	s = s + s
	for i := 0; i < n; i += g {
		cArr := []byte(s[i : i+n])
		addChar(cArr, a, 1)
		if oddB {
			addChar(cArr, a, 0)
		}
		cand := string(cArr)
		if cand < ans {
			ans = cand
		}
	}
	return ans
}
