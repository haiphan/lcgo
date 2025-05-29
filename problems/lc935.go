package problems

// identity matrix for 10x10
var id10 = [][]int{
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
}

func mtMul(a, b [][]int) [][]int {
	n := len(a)
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, n)
		for j := range n {
			for k := range n {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

func mtExp(mat [][]int, k int) [][]int {
	res := id10
	for k > 0 {
		if k%2 == 1 {
			res = mtMul(res, mat)
		}
		mat = mtMul(mat, mat)
		k /= 2
	}
	return res
}

func knightDialer(n int) int {
	mat := [][]int{
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
	}
	res := mtExp(mat, n-1)
	sum := 0
	for i := range 10 {
		for j := range 10 {
			sum = (sum + res[i][j]) % 1000000007
		}
	}
	return sum
}
