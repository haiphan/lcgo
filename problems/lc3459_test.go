package problems

import (
	"reflect"
	"testing"
)

func TestLenOfVDiagonal(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}},
			want: 5,
		},
		{
			name: "Example 2",
			grid: [][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lenOfVDiagonal(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lenOfVDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
