package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestSubtreeWithAllDeepest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			name: "Example 1",
			root: utils.DeserializeTree("3,5,6,#,#,2,7,#,#,4,#,#,1,0,#,#,8,#,#"),
			want: "2,7,#,#,4,#,#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sub := subtreeWithAllDeepest(tt.root)
			got := utils.SerializeTree(sub)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtreeWithAllDeepest() = %v, want %v", got, tt.want)
			}
		})
	}
}
