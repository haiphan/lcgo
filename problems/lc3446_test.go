package problems

import (
	"reflect"
	"testing"
)

func TestSortMatrix(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}},
			want: [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}},
		},
		{
			name: "Example 2",
			grid: [][]int{{0, 1}, {1, 2}},
			want: [][]int{{2, 1}, {1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortMatrix(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
