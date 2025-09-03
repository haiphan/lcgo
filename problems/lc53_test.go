package problems

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
