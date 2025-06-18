package problems

func rob3(root *TreeNode) int {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		withRoot := node.Val + l[1] + r[1]
		noRoot := max(l[0], l[1]) + max(r[0], r[1])
		l[0], l[1] = withRoot, noRoot
		return l
	}
	ans := dfs(root)
	return max(ans[0], ans[1])
}
