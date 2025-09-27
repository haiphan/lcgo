package problems

import (
	"reflect"
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   float64
	}{
		{
			name:   "Example 1",
			points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			want:   2.0,
		},
		{
			name:   "Example 2",
			points: [][]int{{1, 0}, {0, 0}, {0, 1}},
			want:   0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestTriangleArea(tt.points)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
