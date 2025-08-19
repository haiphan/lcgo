package problems

import (
	"reflect"
	"testing"
)

func TestZeroFilledSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 0, 0, 2, 0, 0, 4},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{0, 0, 0, 2, 0, 0},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zeroFilledSubarray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroFilledSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
