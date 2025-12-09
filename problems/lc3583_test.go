package problems

import (
	"reflect"
	"testing"
)

func TestSpecialTriplets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{6, 3, 6},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 0, 0},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := specialTriplets(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("specialTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
