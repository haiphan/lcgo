package problems

var MAXV int = 1e5
var FACT []int = make([]int, MAXV)
var IFACT []int = make([]int, MAXV)

func initCache() {
	if FACT[0] != 0 {
		return
	}
	FACT[0] = 1
	for i := 1; i < MAXV; i++ {
		FACT[i] = i * FACT[i-1] % MOD
	}
	IFACT[MAXV-1] = pow(FACT[MAXV-1], MOD-2)
	for i := MAXV - 1; i > 0; i-- {
		IFACT[i-1] = (IFACT[i] * i) % MOD
	}
}

func fact(n int) int {
	return FACT[n]
}

func pow(n, p int) int {
	res := 1
	for p > 0 {
		if p%2 == 1 {
			res = (res * n) % MOD
		}
		n = (n * n) % MOD
		p = p >> 1
	}
	return res
}

func invFact(n int) int {
	return IFACT[n]
}

func comb(n, k int) int {
	return fact(n) * invFact(n-k) % MOD * invFact(k) % MOD
}

func countGoodArrays(n int, m int, k int) int {
	initCache()
	return m * pow(m-1, n-k-1) % MOD * comb(n-1, k) % MOD
}
