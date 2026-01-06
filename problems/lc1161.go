package problems

func maxLevelSum(root *TreeNode) int {
	sl := make([]int, 0, 10000)
	var dfs func(node *TreeNode, i int)
	dfs = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		if i == len(sl) {
			sl = append(sl, 0)
		}
		sl[i] += node.Val
		dfs(node.Left, i+1)
		dfs(node.Right, i+1)
	}
	dfs(root, 0)
	maxs := sl[0]
	for _, x := range sl {
		maxs = max(maxs, x)
	}
	for i, x := range sl {
		if maxs == x {
			return i + 1
		}
	}
	return 1
}
