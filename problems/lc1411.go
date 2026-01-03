package problems

func numOfWays(n int) int {
	// A: patterns with different colors at the end (ABC)
	// B: patterns with same colors at the end (ABA)
	const MOD = 1e9 + 7
	A, B := 6, 6
	for i := 2; i <= n; i++ {
		newA := (2*A + 2*B) % MOD
		newB := (newA + B) % MOD
		A = newA
		B = newB
	}

	return (A + B) % MOD
}
