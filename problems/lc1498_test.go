package problems

import (
	"reflect"
	"testing"
)

func TestNumSubseq(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{3, 5, 6, 7},
			target: 9,
			want:   4,
		},
		{
			name:   "Example 2",
			nums:   []int{3, 3, 6, 8},
			target: 10,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSubseq(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
