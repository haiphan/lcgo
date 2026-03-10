package problems

import "fmt"

type Lc3129 struct {
	fact    [100001]int
	invFact [100001]int
	nMax    int
}

var lc3129 Lc3129 = Lc3129{
	fact:    [100001]int{},
	invFact: [100001]int{},
	nMax:    -1,
}

func fillData3129(n int) {
	if lc3129.nMax >= n {
		return
	}
	lc3129.fact[0] = 1
	// start from lc3129.nMax + 1 to n
	start := max(lc3129.nMax+1, 1)
	for i := start; i <= n; i++ {
		lc3129.fact[i] = (lc3129.fact[i-1] * i) % MOD
	}
	lc3129.invFact[n] = powMod(lc3129.fact[n], MOD-2)
	for i := n - 1; i >= lc3129.nMax+1; i-- {
		lc3129.invFact[i] = (lc3129.invFact[i+1] * (i + 1)) % MOD
	}
	lc3129.nMax = n
}

func calcComb3129(n, k int) int {
	if k > n || k < 0 {
		return 0
	}
	return (lc3129.fact[n] * lc3129.invFact[k] % MOD) * lc3129.invFact[n-k] % MOD
}

// count the number of ways to pick k groups of l elements from n elements.
func pickNKL(n, k, l int) int {
	if k > n || k < 0 {
		return 0
	}
	ans := 0
	maxj := (n - k) / l
	for j := 0; j <= maxj; j++ {
		ways := calcComb3129(n-1-j*l, k-1) * calcComb3129(k, j) % MOD
		if j%2 == 0 {
			ans = (ans + ways) % MOD
		} else {
			ans = (ans - ways + MOD) % MOD
		}
	}
	return ans
}

func numberOfStableArrays(zero int, one int, limit int) int {
	n := zero + one
	fillData3129(n)
	fmt.Println("fill", lc3129.nMax, lc3129.fact[:10], lc3129.invFact[:10])
	maxK := min(zero, one+1)
	one_cache := make([]int, maxK+2)
	for k := 1; k <= maxK+1; k++ {
		one_cache[k] = pickNKL(one, k, limit)
	}
	ans := 0
	for k := 1; k <= maxK; k++ {
		c0 := pickNKL(zero, k, limit)
		if c0 == 0 {
			continue
		}
		c1 := (one_cache[k-1] + 2*one_cache[k] + one_cache[k+1]) % MOD
		ans = (ans + c0*c1) % MOD
	}
	return ans
}
