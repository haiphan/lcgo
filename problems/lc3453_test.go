package problems

import (
	"math"
	"reflect"
	"testing"
)

func TestSeparateSquares(t *testing.T) {
	tests := []struct {
		name    string
		squares [][]int
		want    float64
	}{
		{
			name:    "Example 1",
			squares: [][]int{{0, 0, 2}, {1, 1, 1}},
			want:    1.16667,
		},
	}
	const EPSILON float64 = 1e-5
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := separateSquares(tt.squares)
			if !reflect.DeepEqual(true, math.Abs(got-tt.want) < EPSILON) {
				t.Errorf("separateSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
