package problems

import (
	"reflect"
	"testing"
)

func TestCountSquares(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "Example 1",
			matrix: [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}},
			want:   15,
		},
		{
			name:   "Example 2",
			matrix: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSquares(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
