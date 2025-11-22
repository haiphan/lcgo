package problems

import (
	"reflect"
	"testing"
)

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{3, 6, 9},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumOperations(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
