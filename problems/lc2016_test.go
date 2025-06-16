package problems

import (
	"reflect"
	"testing"
)

func TestMaximumDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{7, 1, 5, 4},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{9, 4, 3, 2},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumDifference(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
