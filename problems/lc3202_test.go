package problems

import (
	"reflect"
	"testing"
)

func TestMaximumLength2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{1, 4, 2, 3, 1, 4},
			k:    3,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumLength2(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumLength2() = %v, want %v", got, tt.want)
			}
		})
	}
}
