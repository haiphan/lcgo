package problems

import (
	"reflect"
	"testing"
)

func TestMinimumDifference1984(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{90},
			k:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDifference1984(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumDifference1984() = %v, want %v", got, tt.want)
			}
		})
	}
}
