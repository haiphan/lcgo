package problems

import (
	"reflect"
	"testing"
)

func TestFindNumberOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 5, 4, 7},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 2, 2, 2},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNumberOfLIS(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
