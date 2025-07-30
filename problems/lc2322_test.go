package problems

import (
	"reflect"
	"testing"
)

func TestMinimumScore(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		edges [][]int
		want  int
	}{
		{
			name:  "Example 1",
			nums:  []int{1, 5, 5, 4, 11},
			edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
			want:  9,
		},
		{
			name:  "Example 2",
			nums:  []int{5, 5, 2, 4, 4, 2},
			edges: [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumScore(tt.nums, tt.edges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
