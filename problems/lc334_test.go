package problems

import (
	"reflect"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "Example 2",
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := increasingTriplet(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
