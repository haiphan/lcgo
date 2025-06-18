package problems

import (
	"reflect"
	"testing"
)

func TestDivideArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
			k:    2,
			want: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			name: "Example 2",
			nums: []int{2, 4, 2, 2, 5, 2},
			k:    2,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divideArray(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
