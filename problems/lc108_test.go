package problems

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{
			name: "Example 1",
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			name: "Example 2",
			nums: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)
			want := true
			got := isBalanced(root) && isValidBST(root)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, want)
			}
		})
	}
}
