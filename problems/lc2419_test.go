package problems

import (
	"reflect"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 3, 2, 2},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
