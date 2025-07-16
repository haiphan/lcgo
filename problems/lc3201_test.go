package problems

import (
	"reflect"
	"testing"
)

func TestMaximumLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 1, 1, 2, 1, 2},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumLength(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
