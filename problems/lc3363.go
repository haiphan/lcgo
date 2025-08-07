package problems

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	n1 := n - 1
	total := 0
	for i := range n {
		total += fruits[i][i]
	}
	topDown := func(i, j int) {
		maxPrev := 0
		for d := -1; d <= 1; d++ {
			pj := j + d
			if pj >= n-i && pj < n {
				maxPrev = max(maxPrev, fruits[i-1][pj])
			}
		}
		fruits[i][j] += maxPrev
	}
	for i := 1; i < n1; i++ {
		lower := max(i+1, n-i-1)
		for j := n1; j >= lower; j-- {
			topDown(i, j)
		}
	}
	total += fruits[n-2][n-1]
	leftRight := func(i, j int) {
		maxPrev := 0
		for d := -1; d <= 1; d++ {
			pi := i + d
			if pi >= n-j && pi < n {
				maxPrev = max(maxPrev, fruits[pi][j-1])
			}
		}
		fruits[i][j] += maxPrev
	}
	for j := 1; j < n1; j++ {
		lower := max(j+1, n-j-1)
		for i := n1; i >= lower; i-- {
			leftRight(i, j)
		}
	}
	total += fruits[n-1][n-2]
	return total
}
