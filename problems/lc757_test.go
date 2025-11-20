package problems

import (
	"reflect"
	"testing"
)

func TestIntersectionSizeTwo(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			want:      5,
		},
		{
			name:      "Example 2",
			intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			want:      3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersectionSizeTwo(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
