package problems

const M64 int64 = 1e9 + 7

func countTrapezoids(points [][]int) int {
	pm := make(map[int]int, len(points))
	for _, p := range points {
		pm[p[1]]++
	}
	var sum, cs int64
	for _, c := range pm {
		if c == 1 {
			continue
		}
		cb := int64(c)
		segs := (cb * (cb - 1)) >> 1
		sum += segs
		cs += segs * segs
	}
	ans := (sum*sum - cs) >> 1
	ans %= M64
	return int(ans)
}
