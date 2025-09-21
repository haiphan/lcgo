package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 11, 10, 1, 3},
			k:    10,
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 2, 4, 9},
			k:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations1(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations1() = %v, want %v", got, tt.want)
			}
		})
	}
}
