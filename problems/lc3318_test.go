package problems

import (
	"reflect"
	"testing"
)

func TestFindXSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k, x int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:    6,
			x:    2,
			want: []int{6, 10, 12},
		},
		{
			name: "Example 2",
			nums: []int{3, 8, 7, 8, 7, 5},
			k:    2,
			x:    2,
			want: []int{11, 15, 15, 15, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findXSum(tt.nums, tt.k, tt.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findXSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
