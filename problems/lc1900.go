package problems

const MAXPLAYERS int = 28
const MS int = 784
const CACHE_SIZE int = 21952

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	seen := make([]bool, CACHE_SIZE)
	p1, p2 := firstPlayer-1, secondPlayer-1
	lCnt, mCnt, rCnt := p1, p2-p1-1, n-1-p2
	minv, maxv := n, 0
	mask := (1 << n) - 1
	var dfs func(mask, l, r, round, lCnt, mCnt, rCnt int)
	dfs = func(mask, l, r, round, lCnt, mCnt, rCnt int) {
		if l >= r {
			dfs(mask, 0, n-1, round+1, lCnt, mCnt, rCnt)
		} else if (mask & (1 << l)) == 0 {
			dfs(mask, l+1, r, round, lCnt, mCnt, rCnt)
		} else if (mask & (1 << r)) == 0 {
			dfs(mask, l, r-1, round, lCnt, mCnt, rCnt)
		} else if l == p1 && r == p2 {
			minv = min(minv, round)
			maxv = max(maxv, round)
		} else {
			key := lCnt*MS + mCnt*MAXPLAYERS + rCnt
			if seen[key] {
				return
			}
			seen[key] = true
			if l != p1 && l != p2 {
				nl, nm, nr := lCnt, mCnt, rCnt
				if l < p1 {
					nl--
				}
				if p1 < l && l < p2 {
					nm--
				}
				if p2 < l {
					nr--
				}
				dfs(mask^(1<<l), l+1, r-1, round, nl, nm, nr)
			}

			if r != p1 && r != p2 {
				nl, nm, nr := lCnt, mCnt, rCnt
				if r < p1 {
					nl--
				}
				if p1 < r && r < p2 {
					nm--
				}
				if p2 < r {
					nr--
				}
				dfs(mask^(1<<r), l+1, r-1, round, nl, nm, nr)
			}
		}
	}
	dfs(mask, 0, n-1, 1, lCnt, mCnt, rCnt)
	return []int{minv, maxv}
}
