package problems

import (
	"reflect"
	"testing"
)

func TestMinOperations3512(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		k, want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 9, 7},
			k:    5,
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{4, 1, 3},
			k:    4,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations3512(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations3512() = %v, want %v", got, tt.want)
			}
		})
	}
}
