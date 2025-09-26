package problems

import (
	"reflect"
	"testing"
)

func TestTriangleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 2, 3, 4},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{4, 2, 3, 4},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := triangleNumber(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
