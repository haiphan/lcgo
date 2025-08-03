package problems

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	sp := startPos
	N := len(fruits)
	l, cur, res := 0, 0, 0
	getDist := func(lp, rp int) int {
		if sp >= rp {
			return sp - lp
		}
		if sp <= lp {
			return rp - sp
		}
		return min(sp+rp-2*lp, 2*rp-sp-lp)
	}

	for r := range N {
		cur += fruits[r][1]
		for l <= r && getDist(fruits[l][0], fruits[r][0]) > k {
			cur -= fruits[l][1]
			l++
		}
		res = max(res, cur)
	}
	return res
}
