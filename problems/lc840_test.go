package problems

import (
	"reflect"
	"testing"
)

func TestNumMagicSquaresInside(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{4, 3, 8, 4},
				{9, 5, 1, 9},
				{2, 7, 6, 2},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numMagicSquaresInside(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMagicSquaresInside() = %v, want %v", got, tt.want)
			}
		})
	}
}
