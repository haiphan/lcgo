package problems

import (
	"reflect"
	"testing"
)

func TestMaxKDivisibleComponents(t *testing.T) {
	tests := []struct {
		name   string
		n, k   int
		edges  [][]int
		values []int
		want   int
	}{
		{
			name:   "Example 1",
			n:      5,
			edges:  [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
			values: []int{1, 8, 1, 4, 4},
			k:      6,
			want:   2,
		},
		{
			name:   "Example 2",
			n:      7,
			edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			values: []int{3, 0, 6, 1, 5, 2, 1},
			k:      3,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxKDivisibleComponents(tt.n, tt.edges, tt.values, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxKDivisibleComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
