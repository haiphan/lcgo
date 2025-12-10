package problems

func factMod(n int) int {
	v := 1
	for i := 2; i <= n; i++ {
		v = v * i % MOD
	}
	return v
}
func countPermutations(complexity []int) int {
	n := len(complexity)
	if n == 1 {
		return 1
	}
	minv := complexity[0]
	for i := 1; i < n; i++ {
		if complexity[i] <= minv {
			return 0
		}
	}
	return factMod(n - 1)
}
