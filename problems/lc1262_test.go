package problems

import (
	"reflect"
	"testing"
)

func TestMaxSumDivThree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 6, 5, 1, 8},
			want: 18,
		},
		{
			name: "Example 2",
			nums: []int{4},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSumDivThree(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumDivThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
