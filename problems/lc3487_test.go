package problems

import (
	"reflect"
	"testing"
)

func TestMaxSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 0, 1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
