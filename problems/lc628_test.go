package problems

import (
	"reflect"
	"testing"
)

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			want: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumProduct(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
