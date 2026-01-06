package problems

import (
	"lcgo/utils"
	"reflect"
	"testing"
)

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "Example 1",
			root: utils.DeserializeTree("1,7,7,#,#-8,#,#,0,#,#"),
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxLevelSum(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
