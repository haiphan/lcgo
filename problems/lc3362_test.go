package problems

import (
	"reflect"
	"testing"
)

func TestMaxRemoval(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    int
	}{
		{
			name:    "Example 1",
			nums:    []int{2, 0, 2},
			queries: [][]int{{0, 2}, {0, 2}, {1, 1}},
			want:    1,
		},
		{
			name:    "Example 2",
			nums:    []int{1, 1, 1, 1},
			queries: [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}},
			want:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxRemoval(tt.nums, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
