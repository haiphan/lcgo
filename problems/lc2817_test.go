package problems

import (
	"reflect"
	"testing"
)

func TestMinAbsoluteDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		x    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			x:    3,
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{5, 14, 81},
			x:    2,
			want: 76,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minAbsoluteDifference(tt.nums, tt.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAbsoluteDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
