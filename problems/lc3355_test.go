package problems

import (
	"reflect"
	"testing"
)

func TestIsZeroArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    bool
	}{
		{
			name:    "Example 1",
			nums:    []int{1, 0, 1},
			queries: [][]int{{0, 2}},
			want:    true,
		},
		{
			name:    "Example 2",
			nums:    []int{4, 3, 2, 1},
			queries: [][]int{{1, 3}, {0, 2}},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isZeroArray(tt.nums, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
