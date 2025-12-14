package problems

func numberOfWays2147(corridor string) int {
	cc := corridor
	if len(cc) < 2 {
		return 0
	}
	p := 0
	cnt := 0
	ans := 1
	for i := range cc {
		if cc[i] == 'S' {
			cnt++
			if (cnt&1) == 1 && (cnt > 2) {
				ans = ans * (i - p) % MOD
			}
			p = i
		}
	}
	if cnt < 2 || (cnt&1) == 1 {
		return 0
	}
	return ans
}
