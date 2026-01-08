package problems

import (
	"reflect"
	"testing"
)

func TestMaxDotProduct(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		want         int
	}{
		{
			name:  "Example 1",
			nums1: []int{2, 1, -2, 5},
			nums2: []int{3, 0, -6},
			want:  18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDotProduct(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
