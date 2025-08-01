package problems

import (
	"reflect"
	"testing"
)

func TestCountMaxOrSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 2},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countMaxOrSubsets(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
