package problems

import (
	"reflect"
	"testing"
)

func TestFindPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, 4, 1, 5},
			k:    2,
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPairs(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
