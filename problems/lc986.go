package problems

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0, len(firstList)+len(firstList))
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][1] < secondList[j][0] {
			i++
		} else if secondList[j][1] < firstList[i][0] {
			j++
		} else {
			res = append(res, []int{max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])})
			if firstList[i][1] < secondList[j][1] {
				i++
			} else {
				j++
			}
		}
	}
	return res
}
