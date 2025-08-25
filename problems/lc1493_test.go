package problems

import (
	"reflect"
	"testing"
)

func TestLongestSubarrayOfOne(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 0, 1},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarrayOfOne(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestSubarrayOfOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
