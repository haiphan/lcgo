package problems

func productQueries(n int, queries [][]int) []int {
	pw := make([]int, 0)
	x := 1
	for n > 0 {
		if (n & 1) == 1 {
			pw = append(pw, x)
		}
		n = n >> 1
		x = x << 1
	}
	L := len(pw)
	pre := make([]int, L)
	inv := make([]int, L)
	pre[0] = pw[0]
	for i := 1; i < L; i++ {
		pre[i] = pw[i] * pre[i-1] % MOD
	}
	inv[L-1] = intPowMod(pre[L-1], MOD-2, MOD)
	for i := L - 2; i >= 0; i-- {
		inv[i] = pw[i+1] * inv[i+1] % MOD
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		s, e := q[0], q[1]
		if s == e {
			res[i] = pw[s]
			continue
		}
		v := pre[e]
		if s > 0 {
			v = v * inv[s-1] % MOD
		}
		res[i] = v
	}
	return res
}
