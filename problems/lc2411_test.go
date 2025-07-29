package problems

import (
	"reflect"
	"testing"
)

func TestSmallestSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 0, 2, 1, 3},
			want: []int{3, 3, 2, 2, 1},
		},
		{
			name: "Example 2",
			nums: []int{1, 2},
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestSubarrays(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
