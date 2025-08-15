package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations2(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		want  int64
	}{
		{
			name:  "Example 1",
			nums1: []int{0, 0, 1},
			nums2: []int{0, 0, 2},
			k:     0,
			want:  -1,
		},
		{
			name:  "Example 2",
			nums1: []int{4, 3, 1, 4},
			nums2: []int{1, 3, 7, 1},
			k:     3,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations2(tt.nums1, tt.nums2, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations2() = %v, want %v", got, tt.want)
			}
		})
	}
}
