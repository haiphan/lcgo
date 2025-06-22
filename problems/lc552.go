package problems

func checkRecord(n int) int {
	if n == 1 {
		return 3
	}
	// A, L
	l0, l1, l2 := 1, 1, 0    // no A
	al0, al1, al2 := 1, 0, 0 // 1A, 0-2 L
	for i := 1; i < n; i++ {
		p := (l0 + l1 + l2) % MOD        // add P
		a := (al0 + al1 + al2 + p) % MOD // add A
		// add L
		l2, l1 = l1, l0
		al2, al1 = al1, al0
		l0, al0 = p, a
	}
	return (l0 + l1 + l2 + al0 + al1 + al2) % MOD
}
