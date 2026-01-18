package problems

import (
	"reflect"
	"testing"
)

func TestLargestMagicSquare(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestMagicSquare(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestMagicSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
