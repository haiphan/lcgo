package problems

import (
	"reflect"
	"testing"
)

func TestMaxTargetNodes2(t *testing.T) {
	tests := []struct {
		name   string
		edges1 [][]int
		edges2 [][]int
		want   []int
	}{
		{
			name:   "Example 1",
			edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
			edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}},
			want:   []int{8, 7, 7, 8, 8},
		},
		{
			name:   "Example 2",
			edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			edges2: [][]int{{0, 1}, {1, 2}, {2, 3}},
			want:   []int{3, 6, 6, 6, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTargetNodes2(tt.edges1, tt.edges2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTargetNodes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
