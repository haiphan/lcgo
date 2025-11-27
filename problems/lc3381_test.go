package problems

import (
	"reflect"
	"testing"
)

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "Example 1",
			nums: []int{1, 2},
			k:    1,
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{-1, -2, -3, -4, -5},
			k:    4,
			want: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubarraySum(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
