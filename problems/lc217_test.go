package problems

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{2, 7, 11, 15},
			want: false,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
