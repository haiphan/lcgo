package problems

import (
	"reflect"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
