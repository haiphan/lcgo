package problems

import (
	"reflect"
	"testing"
)

func TestFindKDistantIndices(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		key  int
		k    int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{3, 4, 9, 1, 3, 9, 5},
			key:  9,
			k:    1,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 2, 2, 2},
			key:  2,
			k:    2,
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKDistantIndices(tt.nums, tt.key, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
