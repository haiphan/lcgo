package problems

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want []int
	}{
		{
			name: "Example 1",
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "Example 2",
			mat:  [][]int{{1, 2}, {3, 4}},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDiagonalOrder(tt.mat)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
