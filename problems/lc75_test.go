package problems

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Example 2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
