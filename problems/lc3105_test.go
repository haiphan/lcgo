package problems

import (
	"reflect"
	"testing"
)

func TestLongestMonotonicSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 4, 3, 3, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestMonotonicSubarray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
