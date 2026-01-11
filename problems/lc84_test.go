package problems

import (
	"reflect"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "Example 2",
			heights: []int{2, 4},
			want:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
