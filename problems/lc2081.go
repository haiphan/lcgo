package problems

import "strconv"

var PAR []int = []int{1, 0}

func ReverseInt(n int) int {
	reversedNum := 0
	for n > 0 {
		digit := n % 10
		reversedNum = reversedNum*10 + digit
		n /= 10
	}

	return reversedNum
}

func isKPa(x int, k int) bool {
	s := strconv.FormatInt(int64(x), k)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func kMirror(k int, n int) int64 {
	cnt := 0
	res, l := 0, 1
	curLen := 0
	for cnt < n {
		curLen++
		r := l * 10
		for _, p := range PAR {
			if cnt == n {
				break
			}
			for i := l; i < r; i++ {
				pre := i
				if p == 1 {
					pre /= 10
				}
				ri := ReverseInt(i)
				cand := (pre * intPow(10, curLen)) + ri
				if isKPa(cand, k) {
					cnt++
					res += cand
					if cnt == n {
						break
					}
				}
			}
		}
		l = r
	}
	return int64(res)
}
