package problems

type NodeDep struct {
	d    int
	node *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, d int) NodeDep
	dfs = func(node *TreeNode, d int) NodeDep {
		if node == nil {
			return NodeDep{d: d, node: node}
		}
		l, r := dfs(node.Left, d+1), dfs(node.Right, d+1)
		if l.d == r.d {
			return NodeDep{d: l.d, node: node}
		} else if l.d > r.d {
			return l
		} else {
			return r
		}
	}
	ans := dfs(root, 0)
	return ans.node
}
