package problems

import (
	"reflect"
	"testing"
)

func TestMaxDistinctElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 2, 3, 3, 4},
			k:    2,
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{4, 4, 4, 4},
			k:    1,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistinctElements(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDistinctElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
