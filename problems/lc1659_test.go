package problems

import (
	"reflect"
	"testing"
)

func TestMaximumUniqueSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{4, 2, 4, 5, 6},
			want: 17,
		},
		{
			name: "Example 2",
			nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumUniqueSubarray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
