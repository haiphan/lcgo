package problems

import (
	"reflect"
	"testing"
)

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{5, 2, 3, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumPairRemoval(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumPairRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
