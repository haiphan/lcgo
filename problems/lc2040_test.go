package problems

import (
	"reflect"
	"testing"
)

func TestKthSmallestProduct(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int64
		want  int64
	}{
		{
			name:  "Example 1",
			nums1: []int{-4, -2, 0, 3},
			nums2: []int{2, 4},
			k:     6,
			want:  0,
		},
		{
			name:  "Example 2",
			nums1: []int{2, 5},
			nums2: []int{3, 4},
			k:     2,
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallestProduct(tt.nums1, tt.nums2, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
