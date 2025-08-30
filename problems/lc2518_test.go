package problems

import (
	"reflect"
	"testing"
)

func TestCountPartitions(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			k:    4,
			want: 6,
		},
		{
			name: "Example 1",
			nums: []int{6, 6},
			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPartitions(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
