package utils

import (
	"fmt"
	"strconv"
	"strings"
)

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

func DeserializeTree(data string) *TreeNode {
	// Split the string into a list of values
	nodes := strings.Split(data, Separator)

	// Use a helper function with a pointer to the index to track our position in the list
	index := 0
	return buildTree(nodes, &index)
}

// buildTree is a recursive helper that consumes the nodes list sequentially
func buildTree(nodes []string, index *int) *TreeNode {
	// 1. Base Case: Check for the NullMarker or if we've consumed all elements
	if *index >= len(nodes) || nodes[*index] == NullMarker {
		*index++ // Consume the marker/element
		return nil
	}

	// 2. Recursive Step: The current element is the root
	valStr := nodes[*index]
	*index++ // Move to the next element

	val, err := strconv.Atoi(valStr)
	if err != nil {
		// Should not happen with well-formed input
		fmt.Printf("Error converting string to int: %v\n", err)
		return nil
	}

	root := &TreeNode{Val: val}

	// 3. Build Left Subtree (recursively consumes the next chunk of the list)
	root.Left = buildTree(nodes, index)

	// 4. Build Right Subtree (recursively consumes the next chunk of the list)
	root.Right = buildTree(nodes, index)

	return root
}
