package problems

import (
	"reflect"
	"testing"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, -2, 3, -2},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{5, -3, 5},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubarraySumCircular(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
