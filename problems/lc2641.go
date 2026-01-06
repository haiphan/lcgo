package problems

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	curSum := 0
	q := make([]*TreeNode, 0, 1e5)
	q = append(q, root)
	i := 0
	for i < len(q) {
		end := len(q)
		nextSum := 0
		for ; i < end; i++ {
			cur := q[i]
			cur.Val += curSum
			hasl, hasr := cur.Left != nil, cur.Right != nil
			cs := 0
			if hasl {
				cs += cur.Left.Val
				q = append(q, cur.Left)
			}
			if hasr {
				cs += cur.Right.Val
				q = append(q, cur.Right)
			}
			nextSum += cs
			if hasl {
				cur.Left.Val = -cs
			}
			if hasr {
				cur.Right.Val = -cs
			}
		}
		curSum = nextSum
	}
	return root
}
