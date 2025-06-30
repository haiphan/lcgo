package problems

import (
	"reflect"
	"testing"
)

func TestFindLHS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLHS(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
