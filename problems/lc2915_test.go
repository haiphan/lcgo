package problems

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubsequence(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
			want:   3,
		},
		{
			name:   "Example 2",
			nums:   []int{4, 1, 3, 2, 1, 5},
			target: 7,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubsequence(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lengthOfLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
