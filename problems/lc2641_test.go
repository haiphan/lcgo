package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestReplaceValueInTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			name: "Example 1",
			root: utils.DeserializeTree("5,4,1,#,#,10,#,#,9,#,7,#,#"),
			want: "0,0,7,#,#,7,#,#,0,#,11,#,#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode := replaceValueInTree(tt.root)
			got := utils.SerializeTree(gotNode)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceValueInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
