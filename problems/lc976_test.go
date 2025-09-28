package problems

import (
	"reflect"
	"testing"
)

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 1, 2},
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 1, 10},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestPerimeter(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
