package problems

func countCoveredBuildings(n int, buildings [][]int) int {
	if len(buildings) < 5 {
		return 0
	}
	BIG := n + 1
	rmax := make([]int, BIG)
	rmin := make([]int, BIG)
	cmax := make([]int, BIG)
	cmin := make([]int, BIG)
	for i := range BIG {
		rmin[i] = BIG
		cmin[i] = BIG
	}
	for _, b := range buildings {
		r, c := b[0], b[1]
		rmin[c] = min(rmin[c], r)
		rmax[c] = max(rmax[c], r)
		cmin[r] = min(cmin[r], c)
		cmax[r] = max(cmax[r], c)
	}
	ans := 0
	for _, b := range buildings {
		r, c := b[0], b[1]
		if r > rmin[c] && r < rmax[c] && c > cmin[r] && c < cmax[r] {
			ans++
		}
	}
	return ans
}
