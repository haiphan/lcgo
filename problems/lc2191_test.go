package problems

import (
	"reflect"
	"testing"
)

func TestSortJumbled(t *testing.T) {
	tests := []struct {
		name       string
		mapping    []int
		nums, want []int
	}{
		{
			name:    "Example 1",
			mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6},
			nums:    []int{991, 338, 38},
			want:    []int{338, 38, 991},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortJumbled(tt.mapping, tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortJumbled() = %v, want %v", got, tt.want)
			}
		})
	}
}
