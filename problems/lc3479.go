package problems

func getLeftGte(tree []int, v, cur, start, end int) int {
	if tree[cur] < v {
		return -1
	}
	if start == end {
		tree[cur] = 0
		return start
	}
	m := (start + end) / 2
	l := 2 * cur
	r := l + 1
	res := -1
	if tree[l] >= v {
		res = getLeftGte(tree, v, l, start, m)
	} else {
		res = getLeftGte(tree, v, r, m+1, end)
	}
	tree[cur] = max(tree[l], tree[r])
	return res
}

func buildTree(arr, tree []int, cur, start, end int) {
	if start == end {
		tree[cur] = arr[start]
		return
	}
	m := (start + end) / 2
	l := 2 * cur
	r := l + 1
	buildTree(arr, tree, l, start, m)
	buildTree(arr, tree, r, m+1, end)
	tree[cur] = max(tree[l], tree[r])
}
func getTree(arr []int) []int {
	N := len(arr)
	tree := make([]int, 4*N)
	buildTree(arr, tree, 1, 0, N-1)
	return tree
}

func numOfUnplacedFruits3(fruits []int, baskets []int) int {
	cnt := len(fruits)
	tree := getTree(baskets)
	N1 := cnt - 1
	for _, f := range fruits {
		if getLeftGte(tree, f, 1, 0, N1) != -1 {
			cnt--
		}
	}
	return cnt
}
