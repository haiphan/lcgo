package problems

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	N := len(status)
	boxOpen := make([]bool, N)
	boxHas := make([]bool, N)
	boxDone := make([]bool, N)
	q := make([]int, 0)
	qi := 0
	res := 0
	for i := range N {
		boxOpen[i] = status[i] == 1
	}
	for _, b := range initialBoxes {
		boxHas[b] = true
		if boxOpen[b] {
			boxDone[b] = true
			res += candies[b]
			q = append(q, b)
		}
	}
	for qi < len(q) {
		cur := q[qi]
		qi++
		for _, b := range keys[cur] {
			boxOpen[b] = true
			if !boxDone[b] && boxHas[b] {
				boxDone[b] = true
				res += candies[b]
				q = append(q, b)
			}
		}
		for _, b := range containedBoxes[cur] {
			boxHas[b] = true
			if !boxDone[b] && boxOpen[b] {
				boxDone[b] = true
				res += candies[b]
				q = append(q, b)
			}
		}
	}
	return res
}
