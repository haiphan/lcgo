package problems

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func getNode(grid [][]int, t, l, n int) *Node {
	if n == 1 {
		return &Node{Val: grid[t][l] == 1, IsLeaf: true}
	}
	h := n / 2
	tl := getNode(grid, t, l, h)
	tr := getNode(grid, t, l+h, h)
	bl := getNode(grid, t+h, l, h)
	br := getNode(grid, t+h, l+h, h)
	same := tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val
	isLeaf := tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf
	if same && isLeaf {
		return &Node{Val: tl.Val, IsLeaf: true}
	}
	return &Node{Val: tl.Val, IsLeaf: false, TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
}

func construct(grid [][]int) *Node {
	return getNode(grid, 0, 0, len(grid))
}
