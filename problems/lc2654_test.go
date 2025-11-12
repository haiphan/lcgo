package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations2654(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 6, 3, 4},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{2, 10, 6, 14},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations2654(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations2654() = %v, want %v", got, tt.want)
			}
		})
	}
}
