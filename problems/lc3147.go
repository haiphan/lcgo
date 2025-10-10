package problems

func maximumEnergy(energy []int, k int) int {
	ee := energy
	n := len(ee)
	maxv := ee[n-1]
	for d := range k {
		start := n - 1 - d
		cur := 0
		for i := start; i >= 0; i -= k {
			cur += ee[i]
			maxv = max(maxv, cur)
		}
	}
	return maxv
}
