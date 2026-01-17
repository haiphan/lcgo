package problems

import (
	"reflect"
	"testing"
)

func TestLargestSquareArea(t *testing.T) {
	tests := []struct {
		name                 string
		bottomLeft, topRight [][]int
		want                 int64
	}{
		{
			name:       "Example 1",
			bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}},
			topRight:   [][]int{{3, 3}, {4, 4}, {6, 6}},
			want:       1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestSquareArea(tt.bottomLeft, tt.topRight)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
