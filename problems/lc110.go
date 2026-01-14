package problems

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) int
	res := true
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := dfs(root.Left), dfs(root.Right)
		if abs(l-r) > 1 {
			res = false
		}
		return 1 + max(l, r)
	}
	dfs(root)
	return res
}
