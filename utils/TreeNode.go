package utils

import "fmt"

const NullMarker = "#"
const Separator = ","

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SerializeTree(root *TreeNode) string {
	if root == nil {
		return NullMarker
	}

	// Root -> Left -> Right
	leftSerialized := SerializeTree(root.Left)
	rightSerialized := SerializeTree(root.Right)

	// Combine components into a comma-separated string
	return fmt.Sprintf("%d%s%s%s%s", root.Val, Separator, leftSerialized, Separator, rightSerialized)
}
