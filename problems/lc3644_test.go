package problems

import (
	"reflect"
	"testing"
)

func TestSortPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{0, 3, 2, 1},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 3, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortPermutation(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
