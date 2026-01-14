package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "Example 1",
			root: utils.DeserializeTree("3,9,#,#,20,15,#,#,7,#,#"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
