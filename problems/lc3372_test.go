package problems

import (
	"reflect"
	"testing"
)

func TestMaxTargetNodes(t *testing.T) {
	tests := []struct {
		name   string
		edges1 [][]int
		edges2 [][]int
		k      int
		want   []int
	}{
		{
			name:   "Example 1",
			edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
			edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}},
			k:      2,
			want:   []int{9, 7, 9, 8, 8},
		},
		{
			name:   "Example 2",
			edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			edges2: [][]int{{0, 1}, {1, 2}, {2, 3}},
			k:      1,
			want:   []int{6, 3, 3, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTargetNodes(tt.edges1, tt.edges2, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTargetNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
