package problems

import (
	"reflect"
	"testing"
)

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, 2},
			want: -1,
		},
		{
			name: "Example 2",
			nums: []int{7, 9, 5, 8, 1, 3},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDifference(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
