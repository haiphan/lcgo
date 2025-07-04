package problems

var PAR []bool = []bool{true, false}

func isKPa(x int, k int) bool {
	orig, rev := x, 0
	for x > 0 {
		rev = rev*k + x%k
		x /= k
	}
	return orig == rev
}

func genP10(x int, isOdd bool) int {
	res := x
	if isOdd {
		x /= 10
	}
	for t := x; t > 0; t /= 10 {
		res = res*10 + t%10
	}
	return res
}

func kMirror(k int, n int) int64 {
	cnt := 0
	res, l := 0, 1
	p10 := 1
	for cnt < n {
		p10 *= 10
		r := l * 10
		for _, odd := range PAR {
			if cnt == n {
				break
			}
			for i := l; i < r; i++ {
				cand := genP10(i, odd)
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
