package problems

import (
	"reflect"
	"testing"
)

func TestInterchangeableRectangles(t *testing.T) {
	tests := []struct {
		name       string
		rectangles [][]int
		want       int64
	}{
		{
			name:       "Example 1",
			rectangles: [][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}},
			want:       6,
		},
		{
			name:       "Example 2",
			rectangles: [][]int{{4, 5}, {7, 8}},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := interchangeableRectangles(tt.rectangles)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interchangeableRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
