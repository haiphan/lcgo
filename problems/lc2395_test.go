package problems

import (
	"reflect"
	"testing"
)

func TestFindSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "Example 2",
			nums: []int{4, 2, 4},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubarrays(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
