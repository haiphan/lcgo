package problems

import (
	"reflect"
	"testing"
)

func TestHasIncreasingSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
			k:    5,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasIncreasingSubarrays(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hasIncreasingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
