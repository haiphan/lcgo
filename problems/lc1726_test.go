package problems

import (
	"reflect"
	"testing"
)

func TestTupleSameProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 3, 4, 6},
			want: 8,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 4, 5, 10},
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tupleSameProduct(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
