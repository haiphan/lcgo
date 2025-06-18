package problems

import (
	"reflect"
	"testing"
)

func TestRob3(t *testing.T) {
	tree1 := &TreeNode{Val: 3}
	tree1.Left = &TreeNode{Val: 2}
	tree1.Right = &TreeNode{Val: 3}
	tree1.Left.Right = &TreeNode{Val: 3}
	tree1.Right.Right = &TreeNode{Val: 1}

	tree2 := &TreeNode{Val: 3}
	tree2.Left = &TreeNode{Val: 4}
	tree2.Right = &TreeNode{Val: 5}
	tree2.Left.Left = &TreeNode{Val: 1}
	tree2.Left.Right = &TreeNode{Val: 3}
	tree2.Right.Right = &TreeNode{Val: 1}
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "Example 1",
			root: tree1,
			want: 7,
		},
		{
			name: "Example 1",
			root: tree2,
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob3(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rob3() = %v, want %v", got, tt.want)
			}
		})
	}
}
