package problems

import (
	"reflect"
	"testing"
)

func TestCountPartitions3578(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		k, want int
	}{
		{
			name: "Example 1",
			nums: []int{9, 4, 1, 3, 7},
			k:    4,
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{3, 3, 4},
			k:    0,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPartitions3578(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPartitions3578() = %v, want %v", got, tt.want)
			}
		})
	}
}
